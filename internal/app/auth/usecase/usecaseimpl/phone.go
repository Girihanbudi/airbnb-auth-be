package usecaseimpl

import (
	"airbnb-auth-be/env/appcontext"
	errpreset "airbnb-auth-be/internal/app/auth/preset/error"
	"airbnb-auth-be/internal/app/auth/preset/request"
	"airbnb-auth-be/internal/app/auth/preset/response"
	transutil "airbnb-auth-be/internal/app/translation/util"
	otpcache "airbnb-auth-be/internal/pkg/cache/otp"
	"airbnb-auth-be/internal/pkg/json"
	msgpreset "airbnb-auth-be/internal/pkg/messaging/preset"
	"airbnb-auth-be/internal/pkg/stderror"
	"airbnb-auth-be/internal/pkg/svcuser/transport/rpc"
	userrpc "airbnb-auth-be/internal/pkg/svcuser/transport/rpc"
	"airbnb-auth-be/internal/pkg/util"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (u Usecase) ContinueWithPhone(ctx *gin.Context, cmd request.ContinueWithPhone) (res response.ContinueWithPhone, err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Validate command request
	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	// Get country using request phone code
	getCountryCmd := rpc.GetCountryByPhoneCodeCmd{Code: int32(cmd.CountryCode)}
	if _, getCountryErr := u.SvcUser.Country.GetCountryByPhoneCode(ctx, &getCountryCmd); getCountryErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getCountryErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	user := &userrpc.User{}

	// Update or create user if not exist
	getUserCmd := userrpc.GetUserByPhoneCmd{
		CountryCode: int32(cmd.CountryCode),
		PhoneNumber: cmd.PhoneNumber,
	}
	if recordUser, getUserErr := u.SvcUser.User.GetUserByPhone(ctx, &getUserCmd); getUserErr != nil {
		user.CountryCode = int32(cmd.CountryCode)
		user.PhoneNumber = cmd.PhoneNumber
		user.Role = userrpc.Role_user.String()

		// Create user default setting
		var userDefaultSetting userrpc.UserDefaultSetting
		userDefaultSetting.UserId = user.Id
		userDefaultSetting.Locale = clientLocale
		userDefaultSetting.Currency = appcontext.GetCurrency(ctx)

		user.DefaultSetting = &userDefaultSetting

		// Insert new user to database
		createUserRes, createUserErr := u.SvcUser.User.CreateUser(ctx, user)
		if createUserErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
		user.Id = createUserRes.Id

	} else {
		user = recordUser
	}

	// Create and store generated OTP for further use
	otp, err := u.createAndStoreOtp(ctx, user.Id)
	if err != nil {
		return
	}

	// Trigger an event to send OTP using sms to user phone number
	userPhoneNumber := fmt.Sprintf("+%d%s", cmd.CountryCode, cmd.PhoneNumber)
	recipients := []string{userPhoneNumber}
	template, err := transutil.TranslateMessage(ctx, "otp", clientLocale)
	if err != nil {
		return
	}
	message := fmt.Sprintf(template, otp)
	payload := msgpreset.SendSmsPayload{
		Recipients: recipients,
		Body:       message,
	}
	msg := msgpreset.SendSms{
		Type:    "otp",
		Context: "signin",
		Payload: *json.Set(payload),
	}
	if _, _, produceEventErr := u.EventProducer.ProduceMessage("sms.send.init", msg); produceEventErr != nil {
		err = transutil.TranslateError(ctx, errpreset.EvtSendMsgFailed, clientLocale)
		return
	}

	// Set user verified bool to notify merchant if user need to register or able to sign in
	res.IsVerified = user.VerifiedAt != nil
	return
}

func (u Usecase) CompletePhoneRegistration(ctx *gin.Context, cmd request.CompletePhoneRegistration) (err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Validate command request
	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	// Get user id from cache by provided OTP
	userId, extractOtpErr := otpcache.Get(cmd.Otp)
	if extractOtpErr != nil {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	// Get user object for updating user information
	getUserCmd := userrpc.GetUserCmd{Id: userId}
	user, getUserErr := u.SvcUser.User.GetUser(ctx, &getUserCmd)
	if getUserErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getUserErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	// Reject request if user already registered
	if user.VerifiedAt != nil {
		err = transutil.TranslateError(ctx, errpreset.UscForbidden, clientLocale)
		return
	}

	// Update user information
	user.VerifiedAt = timestamppb.Now()
	user.FirstName = util.Case(cmd.FirstName, util.CaseLower, util.CaseTitle)
	user.FullName = util.Case(cmd.FirstName+" "+cmd.LastName, util.CaseLower, util.CaseTitle)
	user.Email = cmd.Email
	user.DateOfBirth = timestamppb.New(cmd.ConvertedDateOfBirth())
	user.Role = userrpc.Role_user.String()
	if _, saveUserErr := u.SvcUser.User.UpdateUser(ctx, user); saveUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	// Delete old tokens
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	// Create and store user access and refresh tokens in cache
	return u.createAndStoreTokensPair(ctx, user)
}

func (u Usecase) MakePhoneSession(ctx *gin.Context, cmd request.MakePhoneSession) (err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Get user id from cache by provided OTP
	userId, extractOtpErr := otpcache.Get(cmd.Otp)
	if extractOtpErr != nil {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	// Get user object for continuing sign in
	getUserCmd := userrpc.GetUserCmd{Id: userId}
	user, getUserErr := u.SvcUser.User.GetUser(ctx, &getUserCmd)
	if getUserErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getUserErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	// Reject request if user not registered yet
	if user.VerifiedAt == nil {
		err = transutil.TranslateError(ctx, errpreset.UscForbidden, clientLocale)
		return
	}

	// Delete old tokens
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	// Create and store user access and refresh tokens in cache
	return u.createAndStoreTokensPair(ctx, user)
}

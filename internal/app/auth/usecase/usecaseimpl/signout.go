package usecaseimpl

import (
	"airbnb-auth-be/env/appcontext"
	errpreset "airbnb-auth-be/internal/app/auth/preset/error"
	"airbnb-auth-be/internal/app/auth/preset/request"
	transutil "airbnb-auth-be/internal/app/translation/util"
	authcache "airbnb-auth-be/internal/pkg/cache/auth"
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

func (u Usecase) SignOut(ctx *gin.Context, cmd request.SignOut) (err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Validate command request
	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}

	// Remove access token and it cookie
	atKey, err := u.extractToken(ctx, cmd.AccessToken)
	if err != nil {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}
	if delAtKeyErr := authcache.Del(atKey); delAtKeyErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}
	ctx.SetCookie(
		appcontext.AccessTokenCode,
		cmd.AccessToken,
		-1,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	// Remove refresh token and it cookie
	rtKey, err := u.extractToken(ctx, cmd.AccessToken)
	if err != nil {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}
	if delRtKeyErr := authcache.Del(rtKey); delRtKeyErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}
	ctx.SetCookie(
		appcontext.RefreshTokenCode,
		cmd.AccessToken,
		-1,
		"/sessions",
		env.CONFIG.Domain,
		true,
		true,
	)

	// Remove login indicator cookie
	ctx.SetCookie(
		appcontext.IsLoggedInCode,
		"false",
		-1,
		"/",
		env.CONFIG.Domain,
		true,
		false,
	)

	return
}

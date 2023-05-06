package usecaseimpl

import (
	"airbnb-auth-be/env/appcontext"
	errpreset "airbnb-auth-be/internal/app/auth/preset/error"
	"airbnb-auth-be/internal/app/auth/preset/request"
	transutil "airbnb-auth-be/internal/app/translation/util"
	authcache "airbnb-auth-be/internal/pkg/cache/auth"
	"airbnb-auth-be/internal/pkg/stderror"
	userrpc "airbnb-auth-be/internal/pkg/svcuser/transport/rpc"

	"github.com/gin-gonic/gin"
)

func (u Usecase) RefreshToken(ctx *gin.Context, cmd request.RefreshToken) (err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Validate command request
	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}

	// Extract JTI key from token
	key, err := u.extractToken(ctx, cmd.Token)
	if err != nil {
		return
	}

	// Get claims metadata from cache
	var claims authcache.DefaultClaims
	claimErr := authcache.Get(key, &claims)
	if claimErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}

	// Get user object
	getUserCmd := userrpc.GetUserCmd{Id: claims.UserID}
	user, getUserErr := u.SvcUser.User.GetUser(ctx, &getUserCmd)
	if getUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknInvalid, clientLocale)
		return
	}

	// Delete old tokens
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	// Create and store user access and refresh tokens in cache
	if err = u.createAndStoreTokensPair(ctx, user); err != nil {
		return
	}

	return
}

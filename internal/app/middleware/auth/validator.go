package auth

import (
	"airbnb-auth-be/env/appcontext"
	errpreset "airbnb-auth-be/internal/app/middleware/preset/error"
	transutil "airbnb-auth-be/internal/app/translation/util"
	authcache "airbnb-auth-be/internal/pkg/cache/auth"
	"airbnb-auth-be/internal/pkg/jwt"
	"airbnb-auth-be/internal/pkg/stderror"
	stdresponse "airbnb-auth-be/internal/pkg/stdresponse/rest"
	"context"

	"github.com/gin-gonic/gin"
)

func GinBindAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		at, err := ctx.Cookie(appcontext.AccessTokenCode)
		if err == nil {
			appcontext.SetFromGinRouter(ctx, appcontext.AccessTokenCode, at)
		}

		ctx.Next()
	}
}

func GqlValidateAccessToken(ctx *context.Context) (err error) {
	accessToken := appcontext.GetAccessToken(*ctx)
	clientLocale := appcontext.GetLocale(*ctx)
	if accessToken == nil {
		err = transutil.TranslateError(*ctx, errpreset.TokenNotFound, clientLocale).Error
		return
	}

	userClaims, validateErr := validateJwtToken(*ctx, *accessToken)
	if validateErr != nil {
		err = validateErr.Error
		return
	}

	appcontext.SetFromDefaultRouter(ctx, appcontext.UserClaims, userClaims)
	return
}

func GinValidateAccessToken(ctx *gin.Context) {
	accessToken := appcontext.GetAccessToken(ctx)
	clientLocale := appcontext.GetLocale(ctx)
	if accessToken == nil {
		err := transutil.TranslateError(ctx, errpreset.TokenNotFound, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	userClaims, err := validateJwtToken(ctx, *accessToken)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	appcontext.SetFromGinRouter(ctx, appcontext.UserClaims, userClaims)

	ctx.Next()
}

func GinValidateNoJwtTokenFound(ctx *gin.Context) {
	accessToken := appcontext.GetAccessToken(ctx)
	clientLocale := appcontext.GetLocale(ctx)
	if accessToken != nil {
		err := transutil.TranslateError(ctx, errpreset.UserAlreadyVerified, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Next()
}

func validateJwtToken(ctx context.Context, accessToken string) (tokenClaims *authcache.DefaultClaims, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)
	tokenMetadata := jwt.ExtractTokenMetadata(accessToken)
	if tokenMetadata == nil {
		err = transutil.TranslateError(ctx, errpreset.TokenNotValid, clientLocale)
		return
	}

	claims := *tokenMetadata
	var cacheData authcache.DefaultClaims
	readCacheErr := authcache.Get(claims["jti"].(string), &cacheData)
	if readCacheErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TokenNotFound, clientLocale)
		return
	}

	tokenClaims = &cacheData

	return
}

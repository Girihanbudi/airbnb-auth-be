package rest

import (
	"airbnb-auth-be/env/appcontext"
	"airbnb-auth-be/internal/app/auth/preset/request"
	_ "airbnb-auth-be/internal/app/auth/preset/response"
	transutil "airbnb-auth-be/internal/app/translation/util"
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/stderror"
	stdresponse "airbnb-auth-be/internal/pkg/stdresponse/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ContinueWithGoogle
//
//	@Summary			Get session by using google oauth. User will be redirected to google sign in page to get credential and redirected back to service if user sign in correctly.
//	@Tags					auth
//	@Router				/sessions/google [get]
func (h Handler) ContinueWithGoogle(ctx *gin.Context) {
	h.Auth.ContinueWithGoogle(ctx)
}

// ContinueWithFacebook
//
//	@Summary			Get session by using facebook oauth. User will be redirected to facebook sign in page to get credential and redirected back to service if user sign in correctly.
//	@Tags					auth
//	@Router				/sessions/facebook [get]
func (h Handler) ContinueWithFacebook(ctx *gin.Context) {
	h.Auth.ContinueWithFacebook(ctx)
}

// ContinueWithPhone
//
//		@Summary			Get session by using user phone number. An OTP code will be sent to user phone.
//		@Tags					auth
//		@Accept				json
//		@Produce			json
//	 	@Param				countryCode			body		int			true		"Country Code"
//	 	@Param				phoneNumber			body		string	true		"Phone Number"
//		@Success			201	{object}	response.ContinueWithPhone
//		@Router				/sessions/phone/initial [post]
func (h Handler) ContinueWithPhone(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)
	var req request.ContinueWithPhone
	if bindErr := ctx.ShouldBindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	res, err := h.Auth.ContinueWithPhone(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	stdresponse.GinMakeHttpResponse(ctx, http.StatusCreated, res, nil)
}

// CompletePhoneRegistration
//
//	@Summary		Complete registration using phone number if not verified yet.
//	@Tags				auth
//	@Accept			json
//	@Param			otp						body		string	true		"OTP"
//	@Param			firstName			body		string	true		"First Name"
//	@Param			lastName			body		string	false		"Last Name"
//	@Param			email					body		string	true		"Email"
//	@Param			dateOfBirth		body		string	true		"Date of Birth"
//	@Router			/sessions/phone/complete [post]
func (h Handler) CompletePhoneRegistration(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)
	var req request.CompletePhoneRegistration
	if bindErr := ctx.BindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	err := h.Auth.CompletePhoneRegistration(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

// MakePhoneSession
//
//	@Summary			Make a session usign phone number if user already verified.
//	@Tags					auth
//	@Accept				json
//	@Param				otp						body		string	true		"OTP"
//	@Router				/sessions/phone/generate [post]
func (h Handler) MakePhoneSession(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)
	var req request.MakePhoneSession
	if bindErr := ctx.BindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	err := h.Auth.MakePhoneSession(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

func (h Handler) OauthGoogleCallback(ctx *gin.Context) {
	err := h.Auth.OauthGoogleCallback(ctx)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

func (h Handler) OauthFacebookCallback(ctx *gin.Context) {
	err := h.Auth.OauthFacebookCallback(ctx)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

// ContinueWithGoogle
//
//	@Summary			Rotate access and refresh tokens with a new one after access token expired.
//	@Tags					auth
//	@Router				/sessions/refresh [get]
func (h Handler) RefreshToken(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)

	// Read refresh token from Cookie
	rt, readCookieErr := ctx.Cookie(appcontext.RefreshTokenCode)
	if readCookieErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	req := request.RefreshToken{Token: rt}
	err := h.Auth.RefreshToken(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	stdresponse.GinMakeHttpResponse(ctx, http.StatusCreated, nil, nil)
}

// ContinueWithGoogle
//
//	@Summary			Delete user session by removing user access and refresh tokens.
//	@Tags					auth
//	@Router				/sessions/signout [get]
func (h Handler) SignOut(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)

	// Read refresh token from Cookie
	at, readCookieErr := ctx.Cookie(appcontext.AccessTokenCode)
	if readCookieErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	// Read refresh token from Cookie
	rt, readCookieErr := ctx.Cookie(appcontext.RefreshTokenCode)
	if readCookieErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	req := request.SignOut{
		AccessToken:  at,
		RefreshToken: rt,
	}
	err := h.Auth.SignOut(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	stdresponse.GinMakeHttpResponse(ctx, http.StatusOK, nil, nil)
}

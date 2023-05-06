package usecase

import (
	"airbnb-auth-be/internal/app/auth/preset/request"
	"airbnb-auth-be/internal/app/auth/preset/response"
	"airbnb-auth-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

type IAuth interface {
	// Get session by using google oauth. User will be redirected to google sign in page to get credential and redirected back to service if user sign in correctly.
	ContinueWithGoogle(ctx *gin.Context)
	// Get session by using facebook oauth. User will be redirected to facebook sign in page to get credential and redirected back to service if user sign in correctly.
	ContinueWithFacebook(ctx *gin.Context)
	// Get session by using user phone number. An OTP code will be sent to user phone.
	ContinueWithPhone(ctx *gin.Context, cmd request.ContinueWithPhone) (res response.ContinueWithPhone, err *stderror.StdError)
	// Complete registration using phone number if not verified yet.
	CompletePhoneRegistration(ctx *gin.Context, cmd request.CompletePhoneRegistration) (err *stderror.StdError)
	// Make a session usign phone number if user already verified.
	MakePhoneSession(ctx *gin.Context, cmd request.MakePhoneSession) (err *stderror.StdError)
	// Handle Google oauth callback after user successfully sign in from Google sign in page.
	OauthGoogleCallback(ctx *gin.Context) (err *stderror.StdError)
	// Handle Facebook oauth callback after user successfully sign in from Facebook sign in page.
	OauthFacebookCallback(ctx *gin.Context) (err *stderror.StdError)
	// Rotate access and refresh tokens with a new one after access token expired.
	RefreshToken(ctx *gin.Context, cmd request.RefreshToken) (err *stderror.StdError)
	// Delete user session by removing user access and refresh tokens.
	SignOut(ctx *gin.Context, cmd request.SignOut) (err *stderror.StdError)
}

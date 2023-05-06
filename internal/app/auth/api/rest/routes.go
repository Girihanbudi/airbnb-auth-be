package rest

import authmid "airbnb-auth-be/internal/app/middleware/auth"

func (h Handler) RegisterApi() {
	sessions := h.Router.Group("/sessions")
	{
		sessions.GET("/google", authmid.GinValidateNoJwtTokenFound, h.ContinueWithGoogle)
		sessions.GET("/facebook", authmid.GinValidateNoJwtTokenFound, h.ContinueWithFacebook)

		phone := sessions.Group("/phone")
		{
			phone.POST("/initial", authmid.GinValidateNoJwtTokenFound, h.ContinueWithPhone)
			phone.POST("/complete", authmid.GinValidateNoJwtTokenFound, h.CompletePhoneRegistration)
			phone.POST("/generate", authmid.GinValidateNoJwtTokenFound, h.MakePhoneSession)
		}

		oauth := sessions.Group("/oauth")
		{
			oauth.GET("/google", h.OauthGoogleCallback)
			oauth.GET("/facebook", h.OauthFacebookCallback)
		}

		sessions.PUT("/refresh", authmid.GinValidateNoJwtTokenFound, h.RefreshToken)
		sessions.DELETE("/signout", h.SignOut)
	}
}

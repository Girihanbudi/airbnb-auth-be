package google

import (
	"airbnb-auth-be/internal/pkg/oauth/google/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Options struct {
	config.Config
}

type Oauth struct {
	Options
	oauth2.Config
	UserInfoApi string
}

func NewGoogleOauth(options Options) Oauth {
	var oauth Oauth
	oauth.ClientID = options.ClientId
	oauth.ClientSecret = options.ClientSecret
	oauth.Endpoint = google.Endpoint
	oauth.UserInfoApi = options.UserInfoApi
	oauth.RedirectURL = options.RedirectUrl
	oauth.Scopes = options.Scopes
	oauth.Options = options

	return oauth
}

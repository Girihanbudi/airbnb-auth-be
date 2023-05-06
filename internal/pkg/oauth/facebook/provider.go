package facebook

import (
	"airbnb-auth-be/internal/pkg/oauth/facebook/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Options struct {
	config.Config
}

type Oauth struct {
	Options
	oauth2.Config
	UserInfoApi string
}

func NewFacebookOauth(options Options) Oauth {
	var oauth Oauth
	oauth.ClientID = options.ClientId
	oauth.ClientSecret = options.ClientSecret
	oauth.Endpoint = facebook.Endpoint
	oauth.UserInfoApi = options.UserInfoApi
	oauth.RedirectURL = options.RedirectUrl
	oauth.Scopes = options.Scopes
	oauth.Options = options

	return oauth
}

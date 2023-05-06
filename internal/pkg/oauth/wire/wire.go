package wire

import (
	"airbnb-auth-be/internal/pkg/oauth/facebook"
	"airbnb-auth-be/internal/pkg/oauth/google"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(google.Options), "*"),
	google.NewGoogleOauth,

	wire.Struct(new(facebook.Options), "*"),
	facebook.NewFacebookOauth,
)

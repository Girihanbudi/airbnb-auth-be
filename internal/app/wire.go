//go:build wireinject
// +build wireinject

package app

import (
	creds "airbnb-auth-be/internal/pkg/credential/wire"
	env "airbnb-auth-be/internal/pkg/env/wire"
	gorm "airbnb-auth-be/internal/pkg/gorm/wire"
	http "airbnb-auth-be/internal/pkg/http/server/wire"
	kafka "airbnb-auth-be/internal/pkg/kafka/wire"
	oauth "airbnb-auth-be/internal/pkg/oauth/wire"

	account "airbnb-auth-be/internal/app/account/wire"
	auth "airbnb-auth-be/internal/app/auth/wire"
	translation "airbnb-auth-be/internal/app/translation/wire"

	svcuser "airbnb-auth-be/internal/pkg/svcuser/wire"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(Options), "*"),
	wire.Struct(new(App), "*"),
)

func NewApp() (*App, error) {
	panic(
		wire.Build(
			env.PackageSet,
			creds.PackageSet,
			gorm.PackageSet,
			http.PackageSet,
			kafka.PackageSet,
			oauth.PackageSet,

			svcuser.PackageSet,

			AppSet,

			account.ModuleSet,
			auth.ModuleSet,
			translation.ModuleSet,
		),
	)
}

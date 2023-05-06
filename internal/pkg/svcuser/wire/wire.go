package wire

import (
	"airbnb-auth-be/internal/pkg/svcuser"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(svcuser.Options), "*"),
	svcuser.NewClient,
)

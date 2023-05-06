package wire

import (
	"airbnb-auth-be/internal/pkg/gorm"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(gorm.Options), "*"),
	gorm.NewORM,
)

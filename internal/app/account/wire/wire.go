package wire

import (
	"airbnb-auth-be/internal/app/account/repo"
	"airbnb-auth-be/internal/app/account/repo/repoimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewAccountRepo,
	wire.Bind(new(repo.IAccount), new(*repoimpl.Repo)),
)

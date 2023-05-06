package wire

import (
	"airbnb-auth-be/internal/app/translation/repo"
	"airbnb-auth-be/internal/app/translation/repo/repoimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewTranslationRepo,
	wire.Bind(new(repo.ITranslation), new(*repoimpl.Repo)),
)

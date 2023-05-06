package repo

import (
	module "airbnb-auth-be/internal/app/translation"
	"context"
)

type ITranslation interface {
	GetErrTranslation(ctx context.Context, code, localeCode string) (translation *module.ErrTranslation, err error)
	GetMsgTranslation(ctx context.Context, code, localeCode string) (translation *module.MsgTranslation, err error)
}

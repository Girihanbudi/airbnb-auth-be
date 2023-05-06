package repoimpl

import (
	module "airbnb-auth-be/internal/app/translation"
	"context"
)

func (r Repo) GetMsgTranslation(ctx context.Context, code, localeCode string) (translation *module.MsgTranslation, err error) {
	err = r.Gorm.DB.
		Where("code = ?", code).
		Where("locale_code = ?", localeCode).
		First(&translation).Error

	return
}

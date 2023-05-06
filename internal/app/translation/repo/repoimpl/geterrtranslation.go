package repoimpl

import (
	module "airbnb-auth-be/internal/app/translation"
	"context"
)

func (r Repo) GetErrTranslation(ctx context.Context, code, localeCode string) (translation *module.ErrTranslation, err error) {
	err = r.Gorm.DB.
		Where("code = ?", code).
		Where("locale_code = ?", localeCode).
		First(&translation).Error

	return
}

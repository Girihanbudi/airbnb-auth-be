package repoimpl

import (
	module "airbnb-auth-be/internal/app/account"
	"context"
)

func (r Repo) CreateOrUpdateAccount(ctx context.Context, account *module.Account) (err error) {
	err = r.Gorm.DB.Save(account).Error
	return
}

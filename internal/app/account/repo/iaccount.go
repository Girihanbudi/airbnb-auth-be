package repo

import (
	module "airbnb-auth-be/internal/app/account"
	"context"
)

type IAccount interface {
	CreateOrUpdateAccount(ctx context.Context, account *module.Account) (err error)
}

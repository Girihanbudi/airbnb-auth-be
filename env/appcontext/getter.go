package appcontext

import (
	"airbnb-auth-be/internal/pkg/cache/auth"
	"context"
)

func GetLocale(ctx context.Context) string {
	value := ctx.Value(LocaleCode)
	if value == nil {
		return LocaleDefault
	}
	return value.(string)
}

func GetCurrency(ctx context.Context) string {
	value := ctx.Value(CurrencyCode)
	if value == nil {
		return CurrencyDefault
	}
	return value.(string)
}

func GetAccessToken(ctx context.Context) *string {
	value := ctx.Value(AccessTokenCode)
	if value == nil {
		return nil
	}
	token := value.(string)

	return &token
}

func GetUserClaims(ctx context.Context) *auth.DefaultClaims {
	value := ctx.Value(UserClaims)
	if value == nil {
		return nil
	}
	userClaims := value.(*auth.DefaultClaims)

	return userClaims
}

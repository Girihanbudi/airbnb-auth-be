package cookie

import (
	"airbnb-auth-be/env/appcontext"
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func BindCurrency(ctx *gin.Context) {
	currency, err := ctx.Cookie(appcontext.CurrencyCode)
	if err != nil {
		CreateCurrency(ctx, nil)
		SetCurrency(ctx, nil)
		return
	}

	SetCurrency(ctx, currency)
}

func CreateCurrency(ctx *gin.Context, val *string) {
	if val == nil {
		newVal := appcontext.CurrencyDefault
		val = &newVal
	}

	ctx.SetSameSite(http.DefaultSameSite())

	ctx.SetCookie(
		appcontext.CurrencyCode,
		*val,
		appcontext.CurrencyDuration,
		"/",
		env.CONFIG.Domain,
		true,
		false,
	)
}

func SetCurrency(ctx *gin.Context, val interface{}) {
	if val == nil {
		appcontext.SetFromGinRouter(ctx, appcontext.CurrencyCode, appcontext.CurrencyDefault)
	} else {
		appcontext.SetFromGinRouter(ctx, appcontext.CurrencyCode, val)
	}
}

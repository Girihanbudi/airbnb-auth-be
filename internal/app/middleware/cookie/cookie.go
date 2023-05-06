package cookie

import (
	"github.com/gin-gonic/gin"
)

func BindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		BindLocale(ctx)
		BindCurrency(ctx)
		ctx.Next()
	}
}

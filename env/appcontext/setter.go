package appcontext

import (
	"context"

	"github.com/gin-gonic/gin"
)

func SetFromGinRouter(ctx *gin.Context, key any, val any) {
	newCtx := context.WithValue(ctx.Request.Context(), key, val)
	ctx.Request = ctx.Request.WithContext(newCtx)
}

func SetFromDefaultRouter(ctx *context.Context, key any, val any) {
	newCtx := context.WithValue(*ctx, key, val)
	*ctx = newCtx
}

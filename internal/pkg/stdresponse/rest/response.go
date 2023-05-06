package rest

import (
	"airbnb-auth-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Data     interface{}        `json:"data,omitempty"`
	Metadata interface{}        `json:"metadata,omitempty"`
	Error    *stderror.StdError `json:"error,omitempty"`
}

func GinMakeHttpResponse(ctx *gin.Context, code int, result interface{}, metadata interface{}) {
	ctx.JSON(code, ResponseBody{
		Data:     result,
		Metadata: metadata,
	})
}

func GinMakeHttpResponseErr(ctx *gin.Context, err *stderror.StdError) {
	err.Message = err.Error.Error()
	ctx.AbortWithStatusJSON(err.HttpCode, ResponseBody{
		Error: err,
	})
}

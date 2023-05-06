package rest

import (
	"github.com/gin-gonic/gin"

	authusecase "airbnb-auth-be/internal/app/auth/usecase"
)

type Options struct {
	Router *gin.Engine

	Auth authusecase.IAuth
}

type Handler struct {
	Options
}

func NewAuthHandler(options Options) *Handler {
	return &Handler{options}
}

package app

import (
	_ "airbnb-auth-be/docs"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func (a App) registerHttpHandler() {
	// register auth handler
	a.AuthHandler.RegisterApi()
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}

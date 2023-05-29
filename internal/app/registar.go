package app

import (
	_ "airbnb-auth-be/docs"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func (a App) registerHttpHandler() {
	// Register rest handler
	a.AuthHandler.RegisterApi()

	// Register swagger documentation
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}

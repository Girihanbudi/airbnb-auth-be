package app

import (
	elasticmid "airbnb-auth-be/internal/app/middleware/elastic"
	"airbnb-auth-be/internal/pkg/cache/auth"
	"airbnb-auth-be/internal/pkg/cache/otp"
	elastic "airbnb-auth-be/internal/pkg/elasticsearch"
	"airbnb-auth-be/internal/pkg/http/server"
	httprouter "airbnb-auth-be/internal/pkg/http/server/router"
	kafkaconsumer "airbnb-auth-be/internal/pkg/kafka/consumer"
	kafkaproducer "airbnb-auth-be/internal/pkg/kafka/producer"
	"airbnb-auth-be/internal/pkg/log"
	"airbnb-auth-be/internal/pkg/svcuser"
	"context"
	"sync"

	"github.com/gin-gonic/gin"

	authrest "airbnb-auth-be/internal/app/auth/api/rest"
	authmid "airbnb-auth-be/internal/app/middleware/auth"
	cookiemid "airbnb-auth-be/internal/app/middleware/cookie"

	translation "airbnb-auth-be/internal/app/translation/repo"
)

var Instance = "App"

type Options struct {
	HttpServer    *server.Server
	EventListener *kafkaconsumer.Listener
	EventProducer *kafkaproducer.Producer
	SvcUser       *svcuser.Client

	Translation translation.ITranslation
	AuthHandler *authrest.Handler
}

type App struct {
	Options
}

// Run all the modules of the app.
func (a App) Run(ctx context.Context) {
	a.runModules(ctx)
	a.stopModules()
}

func (a App) runModules(ctx context.Context) {
	log.Event(Instance, "Starting service and connections...")

	// init app cache
	auth.InitAuthCache()
	otp.InitOtpCache()

	// Init elasticsearch client
	elastic.InitElasticSearch()

	// Create required index in elastic
	elasticmid.CreateIndex()

	// recover from panic
	a.HttpServer.Router.Use(gin.Recovery())

	// GIN apply CORS setting
	a.HttpServer.Router.Use(httprouter.DefaultCORSSetting())

	// GIN log request and response to elastic
	a.HttpServer.Router.Use(elasticmid.LogRequestToElastic())

	// GIN bind all cookie
	a.HttpServer.Router.Use(cookiemid.BindAll())

	// GIN bind access token if any
	// bind access token in all route to adapt with graphql endpoint
	a.HttpServer.Router.Use(authmid.GinBindAccessToken())

	// Register all routes
	a.registerHttpHandler()

	go func() {
		a.HttpServer.Start()
	}()

	<-ctx.Done()
}

func (a App) stopModules() {
	log.Event(Instance, "Stoping service and connections...")

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		a.EventProducer.Stop()
	}()

	go func() {
		defer wg.Done()
		a.SvcUser.Stop()
	}()

	go func() {
		defer wg.Done()
		a.HttpServer.Stop()
	}()

	wg.Wait()
	log.Event(Instance, "successfully stopped service and connections")
}

package wire

import (
	"airbnb-auth-be/internal/pkg/kafka"
	"airbnb-auth-be/internal/pkg/kafka/consumer"
	"airbnb-auth-be/internal/pkg/kafka/producer"
	"airbnb-auth-be/internal/pkg/kafka/router"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(router.Options), "*"),
	router.NewRouter,

	wire.Struct(new(kafka.Options), "*"),
	kafka.NewSaramaClient,

	wire.Struct(new(consumer.Options), "*"),
	consumer.NewEventListener,

	wire.Struct(new(producer.Options), "*"),
	producer.NewEventProducer,
)

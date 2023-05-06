package kafka

import (
	"airbnb-auth-be/internal/pkg/kafka/config"
	"airbnb-auth-be/internal/pkg/kafka/router"
	"airbnb-auth-be/internal/pkg/log"

	"github.com/Shopify/sarama"
)

const Instance string = "Sarama Client"

type Options struct {
	config.Config
	Router *router.Router
}

type Client struct {
	Config *sarama.Config
	Options
}

func NewSaramaClient(options Options) *Client {
	log.Event(Instance, "initializing sarama client...")

	sarama.Logger = log.NewLogger(Instance, false)
	version, err := sarama.ParseKafkaVersion(options.Version)
	if err != nil {
		log.Fatal(Instance, "error parsing Kafka version", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version
	config.ClientID = options.ClientId

	return &Client{
		Config:  config,
		Options: options,
	}
}

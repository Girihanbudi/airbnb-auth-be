package consumer

import (
	"airbnb-auth-be/internal/pkg/kafka"
	"airbnb-auth-be/internal/pkg/kafka/consumer/config"
	"airbnb-auth-be/internal/pkg/kafka/router"
	"airbnb-auth-be/internal/pkg/log"
	"fmt"

	"github.com/Shopify/sarama"
)

const Instance string = "Kafka Consumer"

type TopicHandler struct {
	Topic   string
	Handler func()
}

type Options struct {
	config.Config
	Client *kafka.Client
	Router *router.Router
}

type Listener struct {
	client  *kafka.Client
	isReady chan bool

	Consumer sarama.ConsumerGroup
	Options
}

func NewEventListener(options Options) *Listener {
	log.Event(Instance, "initializing kafka listener...")

	if options.Client == nil {
		log.Fatal(Instance, "sarama client not provided", nil)
	}

	switch options.Assigner {
	case "sticky":
		options.Client.Config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundrobin":
		options.Client.Config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		options.Client.Config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		log.Fatal(Instance, fmt.Sprintf("unrecognized consumer group partition assignor: %s", options.Assigner), nil)
	}

	if options.IsUseOldest {
		options.Client.Config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	consumer, err := sarama.NewConsumerGroup(options.Client.Brokers, options.Group, options.Client.Config)
	if err != nil {
		log.Fatal(Instance, "error creating consumer group client: %v", err)
	}

	log.Event(Instance, fmt.Sprintf("ready to consume from %v", options.Client.Brokers))

	return &Listener{
		client:   options.Client,
		isReady:  make(chan bool),
		Consumer: consumer,
		Options:  options,
	}
}

package producer

import (
	"airbnb-auth-be/internal/pkg/kafka"
	"airbnb-auth-be/internal/pkg/log"

	"github.com/Shopify/sarama"
)

const Instance string = "Kafka Producer"

type TopicHandler struct {
	Topic   string
	Handler func()
}

type Options struct {
	Client *kafka.Client
}

type Producer struct {
	Collector sarama.SyncProducer
}

func NewEventProducer(options Options) *Producer {
	log.Event(Instance, "initializing kafka producer...")

	if options.Client == nil {
		log.Fatal(Instance, "sarama client not provided", nil)
	}

	// set tls configuration
	// tlsConfig := createTlsConfiguration()
	// if tlsConfig != nil {
	// 	options.client.Config.Net.TLS.Config = tlsConfig
	// 	options.client.Config.Net.TLS.Enable = true
	// }

	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.
	options.Client.Config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	options.Client.Config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	options.Client.Config.Producer.Return.Successes = true

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.
	producer, err := sarama.NewSyncProducer(options.Client.Brokers, options.Client.Config)
	if err != nil {
		log.Fatal(Instance, "failed to start producer", err)
	}

	return &Producer{
		Collector: producer,
	}
}

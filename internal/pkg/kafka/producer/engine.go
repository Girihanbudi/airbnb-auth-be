package producer

import (
	"airbnb-auth-be/internal/pkg/log"
	"encoding/json"
	"errors"

	"github.com/Shopify/sarama"
)

func (p *Producer) Stop() {
	log.Event(Instance, "shutting down event producer...")

	if err := p.Collector.Close(); err != nil {
		log.Error(Instance, "failed to shutting down event producer", err)
	}

	log.Event(Instance, "event producer has been shutted down")
}

func (p *Producer) ProduceMessage(topic string, value interface{}) (partition int32, offset int64, err error) {
	b, err := json.Marshal(value)
	if err != nil {
		return 0, 0, errors.New("failed to parse value as byte")
	}

	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	return p.Collector.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
	})
}

package router

import (
	"context"

	"github.com/Shopify/sarama"
)

type EventHandler func(ctx context.Context, msg *sarama.ConsumerMessage)

type Handler struct {
	Topic   string
	Handler EventHandler
}

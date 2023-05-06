package consumer

import (
	"airbnb-auth-be/internal/pkg/kafka/router"
	"airbnb-auth-be/internal/pkg/log"
	"context"

	"github.com/Shopify/sarama"
	"github.com/thoas/go-funk"
)

// ref https://github.com/Shopify/sarama/blob/main/examples/consumergroup/main.go

func (l *Listener) Start(ctx context.Context) error {
	var topics []string
	if len(l.Router.Handlers) > 0 {
		topics = funk.Map(l.Router.Handlers, func(handler router.Handler) string {
			return handler.Topic
		}).([]string)
	}

	log.Event(Instance, "starting listener...")

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := l.Consumer.Consume(ctx, topics, l); err != nil {
				log.Fatal(Instance, "error from consumer", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			l.isReady = make(chan bool)
		}
	}()

	<-l.isReady // Await till the consumer has been set up
	log.Event(Instance, "sarama listener up and running")

	return nil
}

func (l *Listener) Stop() error {
	log.Event(Instance, "shutting down event listener...")
	return l.Consumer.Close()
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (l *Listener) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(l.isReady)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (l *Listener) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (l *Listener) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			for _, route := range l.Router.Handlers {
				if message.Topic == route.Topic {
					route.Handler(context.Background(), message)
				}
			}

			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

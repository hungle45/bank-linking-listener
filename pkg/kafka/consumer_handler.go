package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

type ConsumerHandlerFn func(context.Context, *sarama.ConsumerMessage) error

type ConsumerHandler struct {
	TopicHandlers map[string]ConsumerHandlerFn
}

func (c *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok { // check message to prevent panic (ISSUE: https://github.com/IBM/sarama/issues/2477)
				log.Printf("Message channel was closed: topic %s, partition %d, next_offset %d.\n",
					claim.Topic(), claim.Partition(), claim.HighWaterMarkOffset())
				return nil
			}

			err := c.RouteMessage(session.Context(), message)
			if err != nil {
				return err
			}
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}

func (c *ConsumerHandler) RouteMessage(ctx context.Context, message *sarama.ConsumerMessage) error {
	handler, ok := c.TopicHandlers[message.Topic]
	if !ok {
		return nil
	}

	return handler(ctx, message)
}

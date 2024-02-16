package kafka

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

type consumer struct {
	cli     sarama.ConsumerGroup
	topics  []string
	handler sarama.ConsumerGroupHandler
	stop    chan struct{}
	quit    *sync.WaitGroup
}

type Consumer interface {
	Start()
	Stop()
}

func NewConsumer(client sarama.Client, groupID string, topics []string, handler sarama.ConsumerGroupHandler) (Consumer, error) {
	cli, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		return nil, err
	}
	return &consumer{
		cli:    cli,
		topics: topics,
		handler: handler,
		stop:   make(chan struct{}),
		quit:   &sync.WaitGroup{},
	}, nil
}

func (c *consumer) Start() {
	c.quit.Add(1)
	defer c.quit.Done()

	ctx, cancel := context.WithCancel(context.Background())

	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		for {
			if err := c.cli.Consume(ctx, c.topics, c.handler); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-c.stop
	cancel()
	<-doneCh

	if err := c.cli.Close(); err != nil {
		log.Printf("Error closing client err %v", err)
	}
}

func (c *consumer) Stop() {
	close(c.stop)
	c.quit.Wait()
	log.Println("Consumer stopped")
}

package kafka

import "github.com/IBM/sarama"

type producer struct {
	sarama.SyncProducer
	topics []string
}

type Producer interface {
	sarama.SyncProducer
	Topics() []string
}

func NewProducer(client sarama.Client, topics []string) (Producer, error) {
	cli, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	return &producer{
		SyncProducer: cli,
		topics:       topics,
	}, nil
}

func (p *producer) Topics() []string {
	return p.topics
}

package kafka

import (
	"demo/bank-linking-listener/config"

	"github.com/IBM/sarama"
)

func NewClient(cfg *config.Config) sarama.Client {
	brokerAddrs := cfg.Kafka.Brokers
	config := DefaultConfig()
	cli, err := sarama.NewClient(brokerAddrs, config)
	if err != nil {
		panic(err)
	}

	return cli
}

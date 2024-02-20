package kafka

import (
	"demo/bank-linking-listener/config"

	"github.com/IBM/sarama"
)

func NewClient(cfg *config.KafkaConfig) sarama.Client {
	brokerAddrs := cfg.Brokers
	config := DefaultConfig()
	cli, err := sarama.NewClient(brokerAddrs, config)
	if err != nil {
		panic(err)
	}

	return cli
}

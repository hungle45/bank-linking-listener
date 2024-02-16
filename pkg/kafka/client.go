package kafka

import "github.com/IBM/sarama"

func NewClient(brokerAddrs []string) sarama.Client {
	config := DefaultConfig()
	cli, err := sarama.NewClient(brokerAddrs, config)
	if err != nil {
		panic(err)
	}

	return cli
}

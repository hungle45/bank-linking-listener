package server

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal/delivery/consumer"
	"demo/bank-linking-listener/pkg/kafka"
	"log"
)

type consumerJob struct {
	cfg      *config.Config
	consumer kafka.Consumer
}

func NewConsumerJob(cfg *config.Config, controller *consumer.Controller) Server {
	// setup kafka consumer
	BanklinkingLogConfig := cfg.Kafka[config.BankLinkingLog]
	kafkaClient := kafka.NewClient(BanklinkingLogConfig)

	controller.Routes()

	consumer, err := kafka.NewConsumer(kafkaClient, BanklinkingLogConfig.Group, []string{BanklinkingLogConfig.Topic}, controller)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	return &consumerJob{cfg: cfg, consumer: consumer}
}

func (c *consumerJob) Run() {
	c.consumer.Start()
}

func (c *consumerJob) Shutdown() {
	c.consumer.Stop()
}

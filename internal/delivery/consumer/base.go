package consumer

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/pkg/kafka"
)

type Controller struct {
	kafka.ConsumerHandler
	cfg         *config.Config
	bankService service.BankService
}

func NewController(cfg *config.Config, bankService service.BankService) *Controller {
	return &Controller{
		cfg:         cfg,
		bankService: bankService,
	}
}

func (controller *Controller) Routes() {
	controller.TopicHandlers = map[string]kafka.ConsumerHandlerFn{
		controller.cfg.Kafka[config.BankLinkingLog].Topic: controller.HandleBankLinking,
	}
}

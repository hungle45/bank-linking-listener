package consumer

import (
	"context"
	"demo/bank-linking-listener/internal/delivery/consumer/consumer_dto"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/pkg/kafka"
	"encoding/json"
	"errors"
	"log"

	"github.com/IBM/sarama"
)

type bankLinkingConsumer struct {
	kafka.ConsumerHandler
	bankService service.BankService
}

func NewBankLinkingConsumer(bankService service.BankService) *bankLinkingConsumer {
	return &bankLinkingConsumer{bankService: bankService}
}

func (c *bankLinkingConsumer) HandleBankLinking(ctx context.Context, message *sarama.ConsumerMessage) error {
	var req consumer_dto.BankLinkRequest
	err := json.Unmarshal(message.Value, &req)
	if err != nil {
		return errors.New("failed to unmarshal message")
	}

	rerr := c.bankService.LinkBank(ctx, req.UserID, req.BankCode)
	if rerr != nil {
		return errors.New(rerr.Message())
	}

	log.Printf("Bank %s linked for user %v", req.BankCode, req.UserID)
	return nil
}

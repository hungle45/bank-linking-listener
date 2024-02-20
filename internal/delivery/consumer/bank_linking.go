package consumer

import (
	"context"
	"demo/bank-linking-listener/internal/delivery/consumer/consumer_dto"
	"encoding/json"
	"errors"
	"log"

	"github.com/IBM/sarama"
)

func (c *Controller) HandleBankLinking(ctx context.Context, message *sarama.ConsumerMessage) error {
	var req consumer_dto.BankLinkRequest
	err := json.Unmarshal(message.Value, &req)
	if err != nil {
		return errors.New("failed to unmarshal message")
	}

	err = c.bankService.LinkBank(ctx, req.UserID, req.BankCode)
	if err != nil {
		log.Printf("Failed to link bank for user %v: %s", req.UserID, err)
		return errors.New(err.Error())
	}

	log.Printf("Bank %s linked for user %v", req.BankCode, req.UserID)
	return nil
}

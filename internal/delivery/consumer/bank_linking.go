package consumer

import (
	"context"
	"demo/bank-linking-listener/internal/delivery/consumer/consumer_dto"
	"demo/bank-linking-listener/pkg/errorx"
	"encoding/json"
	"errors"
	"log"

	"github.com/IBM/sarama"
)

func (c *Controller) HandleBankLinking(ctx context.Context, message *sarama.ConsumerMessage) error {
	var req consumer_dto.BankLinkRequest
	err := json.Unmarshal(message.Value, &req)
	if err != nil {
		log.Print("[HandleBankLinking] Failed to unmarshal message")
		return nil
	}

	err = c.bankService.LinkBank(ctx, req.UserID, req.BankCode)
	if err != nil {
		log.Printf("[HandleBankLinking] Failed to link bank %v for user %v: %s", req.BankCode, req.UserID, err)
		if err == errorx.ErrorInternal {
			return errors.New(err.Error())
		}
		return nil
	}

	log.Printf("[HandleBankLinking] Bank %s linked for user %v", req.BankCode, req.UserID)
	return nil
}

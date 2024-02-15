package repository

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, entity.Error)
	Create(ctx context.Context, user entity.User) (*entity.User, entity.Error)
}

type BankRepository interface {
	Create(ctx context.Context, bank entity.Bank) (entity.Bank, entity.Error)
	LinkBank(ctx context.Context, userID uint, bankCode string) entity.Error
}
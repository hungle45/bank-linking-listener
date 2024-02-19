package repository

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

//go:generate mockgen -destination=../mocks/mock_repository_$GOFILE -source=$GOFILE -package=mocks

type UserRepository interface {
	GetByID(ctx context.Context, userID uint) (*entity.User, entity.Error)
	GetByEmail(ctx context.Context, email string) (*entity.User, entity.Error)
	Create(ctx context.Context, user entity.User) (*entity.User, entity.Error)
}

type BankRepository interface {
	Create(ctx context.Context, bank entity.Bank) (entity.Bank, entity.Error)
	LinkBank(ctx context.Context, userID uint, bankCode string) entity.Error
	FetchByUserID(ctx context.Context, userID uint) ([]entity.Bank, entity.Error)
}

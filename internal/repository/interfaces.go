package repository

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

//go:generate mockgen -destination=../mocks/mock_repository_$GOFILE -source=$GOFILE -package=mocks

type UserRepository interface {
	GetByID(ctx context.Context, userID uint) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) (*entity.User, error)
}

type BankRepository interface {
	Create(ctx context.Context, bank entity.Bank) (entity.Bank, error)
	LinkBank(ctx context.Context, userID uint, bankCode string) error
	FetchByUserID(ctx context.Context, userID uint) ([]entity.Bank, error)
}

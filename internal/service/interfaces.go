package service

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

//go:generate mockgen -destination=../mocks/mock_service_$GOFILE -source=$GOFILE -package=mocks

type UserService interface {
	SignIn(ctx context.Context, user entity.User) (string, error)
	GetByID(ctx context.Context, id uint) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUserAccount(ctx context.Context, user entity.User) error
	CreateAdminAccount(ctx context.Context, user entity.User) error
	CreateCustomerAccount(ctx context.Context, user entity.User) error
}

type BankService interface {
	GetBankListByUserID(ctx context.Context, userID uint) ([]entity.Bank, error)
	LinkBank(ctx context.Context, userID uint, bankCode string) error
	CreateBank(ctx context.Context, bank entity.Bank) (entity.Bank, error)
}

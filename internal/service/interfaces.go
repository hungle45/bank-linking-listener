package service

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

type UserService interface {
	SignIn(ctx context.Context, user entity.User) (string, entity.Error)
	GetByID(ctx context.Context, id uint) (entity.User, entity.Error)
	GetByEmail(ctx context.Context, email string) (entity.User, entity.Error)
	CreateUserAccount(ctx context.Context, user entity.User) entity.Error
	CreateAdminAccount(ctx context.Context, user entity.User) entity.Error
	CreateCustomerAccount(ctx context.Context, user entity.User) entity.Error
}

type BankService interface {
	// GetBankListByUserID(ctx context.Context, userID uint) ([]entity.Bank, entity.Error)
	LinkBank(ctx context.Context, userID uint, bankCode string) entity.Error
	CreateBank(ctx context.Context, bank entity.Bank) (entity.Bank, entity.Error)
}

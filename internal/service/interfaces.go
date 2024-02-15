package service

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

type UserService interface {
	SignIn(ctx context.Context, user entity.User) (string, entity.Error)
	GetByEmail(ctx context.Context, email string) (entity.User, entity.Error)
	CreateUserAccount(ctx context.Context, user entity.User) entity.Error
	CreateAdminAccount(ctx context.Context, user entity.User) entity.Error
	CreateCustomerAccount(ctx context.Context, user entity.User) entity.Error
}

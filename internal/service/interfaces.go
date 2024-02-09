package service

import (
	"context"
	"demo/bank-linking-listener/internal/service/entity"
)

type UserService interface {
	SignUp(ctx context.Context, user entity.User) entity.Error
	SignIn(ctx context.Context, user entity.User) (string, entity.Error)
}

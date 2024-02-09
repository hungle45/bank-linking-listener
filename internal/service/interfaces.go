package service

import "demo/bank-linking-listener/internal/service/entity"

type UserService interface {
	SignUp(user entity.User) entity.Error
	SignIn(user entity.User) (string, entity.Error)
}

package repository

import "demo/bank-linking-listener/internal/service/entity"

type UserRepository interface {
	GetByEmail(email string) (*entity.User, entity.Error)
	Create(user entity.User) (*entity.User, entity.Error)
}

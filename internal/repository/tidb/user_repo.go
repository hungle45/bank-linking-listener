package tidb

import (
	"context"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/repository/tidb/tidb_dto"
	"demo/bank-linking-listener/internal/service/entity"
)

type userRepository struct {
	user []tidb_dto.UserModel
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, entity.Error) {
	for _, user := range u.user {
		if user.Email == email {
			return user.ToEntity(), nil
		}
	}
	return nil, entity.NewError(entity.ErrorNotFound, "user not found")
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, entity.Error) {
	userModel := tidb_dto.NewUserModel(&user)
	u.user = append(u.user, *userModel)
	return userModel.ToEntity(), nil
}

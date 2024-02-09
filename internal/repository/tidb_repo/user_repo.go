package tidb_repo

import (
	"context"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/repository/tidb_repo/tidb_dto"
	"demo/bank-linking-listener/internal/service/entity"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, entity.Error) {
	userModel := &tidb_dto.UserModel{}
	if err := u.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.NewError(entity.ErrorNotFound, "user not found")
		}
		return nil, entity.NewError(entity.ErrorInternal, err.Error())
	}
	return userModel.ToEntity(), nil
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, entity.Error) {
	userModel := tidb_dto.NewUserModel(&user)
	if err := u.db.Create(&userModel).Error; err != nil {
		if dbErr, ok := err.(*mysql.MySQLError); ok {
			if dbErr.Number == 1062 {
				return nil, entity.NewError(
					entity.ErrorAlreadyExists, fmt.Sprintf("Email %v has been used", user.Email),
				)
			}
		}
		return nil, entity.NewError(entity.ErrorInternal, err.Error())
	}

	return userModel.ToEntity(), nil
}

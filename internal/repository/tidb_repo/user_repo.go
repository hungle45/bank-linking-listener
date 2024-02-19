package tidb_repo

import (
	"context"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/repository/tidb_repo/tidb_dto"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/errorx"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetByID(ctx context.Context, userID uint) (*entity.User, error) {
	userModel := &tidb_dto.UserModel{}
	if err := u.db.First(&userModel, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(int32(codes.NotFound), "User not found")
		}
		return nil, errorx.ErrorInternal
	}
	return userModel.ToEntity(), nil
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	userModel := &tidb_dto.UserModel{}
	if err := u.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(int32(codes.NotFound), "User not found")
		}
		return nil, errorx.ErrorInternal
	}
	return userModel.ToEntity(), nil
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	userModel := tidb_dto.NewUserModel(&user)
	if err := u.db.Create(&userModel).Error; err != nil {
		if dbErr, ok := err.(*mysql.MySQLError); ok {
			if dbErr.Number == 1062 {
				return nil, errorx.New(int32(codes.AlreadyExists), "Email has been used")
			}
		}
		return nil, errorx.ErrorInternal
	}

	return userModel.ToEntity(), nil
}

package tidb_repo

import (
	"context"
	"demo/bank-linking-listener/internal/infrastructure/tidb"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/repository/tidb_repo/tidb_dto"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/errorx"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type userRepository struct {
	cli *gorm.DB
}

func NewUserRepository(db *tidb.TiDB) repository.UserRepository {
	cli := db.GetClient()
	return &userRepository{cli: cli}
}

func (u *userRepository) GetByID(ctx context.Context, userID uint) (*entity.User, error) {
	userModel := &tidb_dto.UserModel{}
	if err := u.cli.First(&userModel, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(int32(codes.NotFound), "User not found")
		}
		return nil, errorx.ErrorInternal
	}
	return userModel.ToEntity(), nil
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	userModel := &tidb_dto.UserModel{}
	if err := u.cli.Where("email = ?", email).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.New(int32(codes.NotFound), "User not found")
		}
		return nil, errorx.ErrorInternal
	}
	return userModel.ToEntity(), nil
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	userModel := tidb_dto.NewUserModel(&user)
	if err := u.cli.Create(&userModel).Error; err != nil {
		if dbErr, ok := err.(*mysql.MySQLError); ok {
			if dbErr.Number == 1062 {
				return nil, errorx.New(int32(codes.AlreadyExists), "Email has been used")
			}
		}
		return nil, errorx.ErrorInternal
	}

	return userModel.ToEntity(), nil
}

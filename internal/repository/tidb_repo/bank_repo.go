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

type bankRepository struct {
	cli *gorm.DB
}

func NewBankRepository(db *tidb.TiDB) repository.BankRepository {
	cli := db.GetClient()
	return &bankRepository{cli: cli}
}

func (r *bankRepository) LinkBank(ctx context.Context, userID uint, bankCode string) error {
	tx := r.cli.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var userModel tidb_dto.UserModel
	if err := tx.Where("id = ?", userID).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return errorx.New(int32(codes.NotFound), "User not found")
		}
		tx.Rollback()
		return errorx.ErrorInternal
	}

	var bankModel tidb_dto.BankModel
	if err := tx.Where("code = ?", bankCode).First(&bankModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return errorx.New(int32(codes.NotFound), "Bank not found")
		}
		tx.Rollback()
		return errorx.ErrorInternal
	}

	if err := tx.Model(&userModel).Association("Banks").Append(&bankModel); err != nil {
		tx.Rollback()
		return errorx.ErrorInternal
	}

	tx.Commit()
	return nil
}

func (r *bankRepository) Create(ctx context.Context, bank entity.Bank) (entity.Bank, error) {
	bankModel := tidb_dto.NewBankModel(&bank)
	if err := r.cli.Create(&bankModel).Error; err != nil {
		if dbErr, ok := err.(*mysql.MySQLError); ok {
			if dbErr.Number == 1062 {
				return entity.Bank{}, errorx.New(int32(codes.AlreadyExists), "Bank code has been used")
			}
		}
		return entity.Bank{}, errorx.ErrorInternal
	}

	return *bankModel.ToEntity(), nil
}

func (r *bankRepository) FetchByUserID(ctx context.Context, userID uint) ([]entity.Bank, error) {
	var userModel tidb_dto.UserModel
	if err := r.cli.Preload("Banks").First(&userModel, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Bank{}, nil
		}
		return []entity.Bank{}, errorx.ErrorInternal
	}

	var banks []entity.Bank
	for _, bankModel := range userModel.Banks {
		banks = append(banks, *bankModel.ToEntity())
	}

	return banks, nil
}

package tidb_dto

import (
	"demo/bank-linking-listener/internal/service/entity"

	"gorm.io/gorm"
)

type BankModel struct {
	gorm.Model
	Code string `gorm:"type:varchar(100);uniqueIndex"`
	Name string `gorm:"type:varchar(100)"`
}

func NewBankModel(bank *entity.Bank) *BankModel {
	return &BankModel{
		Model: gorm.Model{ID: bank.ID},
		Code:  bank.Code,
		Name:  bank.Name,
	}
}

func (b *BankModel) TableName() string {
	return "banks"
}

func (b *BankModel) ToEntity() *entity.Bank {
	return &entity.Bank{
		ID:   b.ID,
		Code: b.Code,
		Name: b.Name,
	}
}

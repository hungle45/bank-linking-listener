package tidb_dto

import (
	"demo/bank-linking-listener/internal/service/entity"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string      `gorm:"type:varchar(100);uniqueIndex"`
	Password string      `gorm:"type:varchar(100)"`
	Role     entity.Role `gorm:"type:varchar(100)"`
}

func NewUserModel(user *entity.User) *UserModel {
	return &UserModel{
		Model:    gorm.Model{ID: user.ID},
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func (u *UserModel) ToUser() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}

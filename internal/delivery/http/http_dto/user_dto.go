package http_dto

import "demo/bank-linking-listener/internal/service/entity"

type UserSignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *UserSignUpRequest) ToEntity() *entity.User {
	return &entity.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

type UserSignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *UserSignInRequest) ToEntity() *entity.User {
	return &entity.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

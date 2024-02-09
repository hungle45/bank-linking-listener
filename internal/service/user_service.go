package service

import (
	"context"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) SignUp(ctx context.Context, user entity.User) entity.Error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.NewError(entity.ErrorInternal, "failed to hash password")
	}
	user.Password = string(hashedPassword)

	_, rerr := s.userRepo.Create(ctx, user)
	if rerr != nil {
		return rerr
	}

	return nil
}

func (s *userService) SignIn(ctx context.Context, user entity.User) (string, entity.Error) {
	res, rerr := s.userRepo.GetByEmail(ctx, user.Email)
	if rerr != nil || bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password)) != nil {
		return "", entity.NewError(entity.ErrorUnauthenticated, "invalid credentials")
	}

	token, err := utils.GenerateToken(res.Email)
	if err != nil {
		return "", entity.NewError(entity.ErrorInternal, "failed to generate token")
	}

	return token, nil
}

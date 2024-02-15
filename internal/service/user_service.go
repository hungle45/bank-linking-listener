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

func (s *userService) createAccount(ctx context.Context, user entity.User) entity.Error {
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

func (s *userService) CreateUserAccount(ctx context.Context, user entity.User) entity.Error {
	user.Role = entity.UserRole
	return s.createAccount(ctx, user)
}
func (s *userService) CreateCustomerAccount(ctx context.Context, user entity.User) entity.Error {
	user.Role = entity.CustomerRole
	return s.createAccount(ctx, user)
}

func (s *userService) CreateAdminAccount(ctx context.Context, admin entity.User) entity.Error {
	rerr := s.createAccount(ctx, admin)
	if rerr != nil && rerr.ErrorType() != entity.ErrorAlreadyExists {
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

func (s *userService) GetByEmail(ctx context.Context, email string) (entity.User, entity.Error) {
	res, rerr := s.userRepo.GetByEmail(ctx, email)
	if rerr != nil {
		return entity.User{}, rerr
	}

	return *res, nil
}

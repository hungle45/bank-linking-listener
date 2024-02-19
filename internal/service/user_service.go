package service

import (
	"context"
	"demo/bank-linking-listener/internal/repository"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/errorx"
	"demo/bank-linking-listener/pkg/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) createAccount(ctx context.Context, user entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errorx.ErrorInternal
	}
	user.Password = string(hashedPassword)

	_, err = s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) CreateUserAccount(ctx context.Context, user entity.User) error {
	user.Role = entity.UserRole
	return s.createAccount(ctx, user)
}
func (s *userService) CreateCustomerAccount(ctx context.Context, user entity.User) error {
	user.Role = entity.CustomerRole
	return s.createAccount(ctx, user)
}

func (s *userService) CreateAdminAccount(ctx context.Context, admin entity.User) error {
	err := s.createAccount(ctx, admin)
	if err != nil && errorx.GetHTTPCode(err) != http.StatusConflict {
		return err
	}

	return nil
}

func (s *userService) SignIn(ctx context.Context, user entity.User) (string, error) {
	res, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password)) != nil {
		return "", errorx.New(int32(codes.Unauthenticated), "invalid credentials")
	}

	token, err := utils.GenerateToken(res.ID)
	if err != nil {
		return "", errorx.ErrorInternal
	}

	return token, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	res, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	return *res, nil
}

func (s *userService) GetByID(ctx context.Context, id uint) (entity.User, error) {
	res, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return *res, nil
}

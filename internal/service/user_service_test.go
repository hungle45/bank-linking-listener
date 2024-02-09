package service_test

import (
	mock_repository "demo/bank-linking-listener/internal/repository/mock"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	userService := service.NewUserService(mockUserRepo)

	t.Run("success", func(t *testing.T) {
		user := entity.User{
			Email:    "user@gmail.com",
			Password: "password",
		}

		mockUserRepo.EXPECT().
			Create(gomock.Cond(func(x any) bool {
				u := x.(entity.User)
				return u.Email == user.Email &&
					bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)) == nil
			})).
			Return(nil, nil)

		err := userService.SignUp(user)
		require.Nil(t, err)
	})
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	userService := service.NewUserService(mockUserRepo)

	t.Run("success", func(t *testing.T) {
		user := entity.User{
			Email:    "user@gmail.com",
			Password: "password",
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		require.Nil(t, err)

		mockUserRepo.EXPECT().
			GetByEmail(user.Email).
			Return(&entity.User{
				Email:    user.Email,
				Password: string(hashedPassword),
			}, nil)

		token, rerr := userService.SignIn(user)
		require.Nil(t, rerr)
		parsedEmail, err := utils.ParseToken(token)
		require.Nil(t, err)
		require.Equal(t, user.Email, parsedEmail)
	})
}

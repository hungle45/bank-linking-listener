package utils

import (
	"demo/bank-linking-listener/config"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(subject uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour * time.Duration(config.TOKEN_EXP)).Unix(),
	})
	return token.SignedString([]byte(config.SECRET_KEY))
}

func ParseToken(token string) (uint, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check if the signing method is HMAC
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return 0, errors.New("token expired")
	}

	return uint(claims["sub"].(float64)), nil
}

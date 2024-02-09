package utils

import (
	"demo/bank-linking-listener/config"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Hour * time.Duration(config.TOKEN_EXP)).Unix(),
	})
	return token.SignedString([]byte(config.SECRET_KEY))
}

func ParseToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check if the signing method is HMAC
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return "", errors.New("token expired")
	}

	return claims["sub"].(string), nil
}

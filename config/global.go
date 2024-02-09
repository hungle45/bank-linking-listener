package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	SECRET_KEY string
	TOKEN_EXP  int
)

func LoadEnv(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("failed to load env file")
	}

	SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY is not set")
	}

	TOKEN_EXP, err = strconv.Atoi(os.Getenv("TOKEN_EXP"))
	if err != nil {
		log.Fatal("TOKEN_EXP is not set")
	}
}

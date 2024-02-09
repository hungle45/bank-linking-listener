package service_test

import (
	"demo/bank-linking-listener/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.LoadEnv("../../.env")
}

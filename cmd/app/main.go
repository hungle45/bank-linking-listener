package main

import (
	"context"
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal/delivery/consumer"
	httpHandler "demo/bank-linking-listener/internal/delivery/http"
	"demo/bank-linking-listener/internal/infrastructure/tidb"
	"demo/bank-linking-listener/internal/repository/tidb_repo"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/kafka"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv("./.env")
	cfg := config.LoadConfig("./config.yml")
	fmt.Println(cfg)

	db := tidb.NewTiDB(cfg)

	// setup repository
	userRepository := tidb_repo.NewUserRepository(db)
	bankRepository := tidb_repo.NewBankRepository(db)

	// setup service
	userService := service.NewUserService(userRepository)
	bankService := service.NewBankService(bankRepository)

	// setup controller
	httpController := httpHandler.NewController(cfg, userService, bankService)
	consumerController := consumer.NewController(cfg, bankService)

	// setup kafka consumer
	kafkaClient := kafka.NewClient(cfg)
	consumerController.Routes()

	consumer, err := kafka.NewConsumer(kafkaClient, "bank-linking-listener", []string{"bank-linking-log"}, consumerController)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	go consumer.Start()
	defer consumer.Stop()

	// init admin
	admin := entity.User{
		Email:    cfg.Server.Admin.Email,
		Password: cfg.Server.Admin.Password,
		Role:     entity.AdminRole,
	}
	if err := userService.CreateAdminAccount(context.Background(), admin); err != nil {
		log.Fatalf("Failed to create admin account: %s", err)
	}

	// setup router
	r := gin.Default()
	httpController.Routes(r)

	s := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server running at %v:%v", cfg.Server.Host, cfg.Server.Port)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

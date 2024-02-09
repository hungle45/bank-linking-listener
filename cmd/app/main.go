package main

import (
	"demo/bank-linking-listener/config"
	httpHandler "demo/bank-linking-listener/internal/delivery/http"
	"demo/bank-linking-listener/internal/delivery/http/route"
	"demo/bank-linking-listener/internal/repository/tidb"
	"demo/bank-linking-listener/internal/service"
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

	// setup repository
	userRepository := tidb.NewUserRepository()

	// setup service
	userService := service.NewUserService(userRepository)

	// setup handler
	userHandler := httpHandler.NewUserHandler(userService)
	bankHandler := httpHandler.NewBankHandler()

	r := gin.Default()

	v1 := r.Group("/v1")
	routerConfig := route.NewRouterConfig(v1, userHandler, bankHandler)
	routerConfig.Setup()

	s := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server running at %v:%v", cfg.Server.Host, cfg.Server.Port)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

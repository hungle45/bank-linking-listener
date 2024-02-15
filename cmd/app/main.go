package main

import (
	"context"
	"demo/bank-linking-listener/config"
	httpHandler "demo/bank-linking-listener/internal/delivery/http"
	"demo/bank-linking-listener/internal/delivery/http/middleware"
	"demo/bank-linking-listener/internal/delivery/http/route"
	"demo/bank-linking-listener/internal/repository/tidb_repo"
	"demo/bank-linking-listener/internal/repository/tidb_repo/tidb_dto"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/internal/service/entity"
	"demo/bank-linking-listener/pkg/tidb"
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

	db := tidb.NewDB(&cfg.Database)
	conn := db.GetConn()

	if err := conn.AutoMigrate(&tidb_dto.UserModel{}); err != nil {
		log.Fatalf("Failed to migrate user model: %s", err)
	}

	if err := conn.AutoMigrate(&tidb_dto.BankModel{}); err != nil {
		log.Fatalf("Failed to migrate bank model: %s", err)
	}

	// setup repository
	userRepository := tidb_repo.NewUserRepository(conn)

	// setup service
	userService := service.NewUserService(userRepository)

	// setup handler
	userHandler := httpHandler.NewUserHandler(userService)
	bankHandler := httpHandler.NewBankHandler()

	// setup middleware
	authMiddleware := middleware.JWTMiddleware()
	roleMiddleware := middleware.RoleMiddleware(userService)

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

	v1 := r.Group("/v1")
	routerConfig := route.RouterConfig{
		Router:         v1,
		UserHandler:    userHandler,
		BankHandler:    bankHandler,
		JWTMiddleware:  authMiddleware,
		RoleMiddleware: roleMiddleware,
	}
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

package server

import (
	"context"
	"demo/bank-linking-listener/config"
	httpHandler "demo/bank-linking-listener/internal/delivery/http"
	"demo/bank-linking-listener/internal/service"
	"demo/bank-linking-listener/internal/service/entity"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	cfg    *config.Config
	server *http.Server
}

func NewHTTPServer(cfg *config.Config, controller *httpHandler.Controller, userService service.UserService) Server {
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
	controller.Routes(r)

	server := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{
		cfg:    cfg,
		server: server,
	}
}

func (s *HTTPServer) Run() {
	log.Printf("Server running at %v:%v", s.cfg.Server.Host, s.cfg.Server.Port)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func (s *HTTPServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %s", err)
	}
}

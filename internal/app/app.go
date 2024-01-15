package app

import (
	"chat/internal/config"
	httpDelivery "chat/internal/delivery/http"
	"chat/internal/integration"
	"chat/internal/repository"
	httpServer "chat/internal/server"
	"chat/internal/service"
	"chat/pkg/db/postgresql"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	postgresDB, err := postgresql.NewDB(cfg.Database.PostgreDSN)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepositories(postgresDB)

	integrations, err := integration.NewIntegrations(cfg)
	if err != nil {
		fmt.Errorf("can't initialize NewRepositories", "err", err.Error())
	}

	services := service.NewServices(repos, integrations)

	handler := httpDelivery.NewHandlerDelivery(logger, services, "chat")

	srv, err := httpServer.NewServer(cfg, handler)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("🔥 Server stopped due error", "err", err.Error())
		} else {
			logger.Printf("✅ Server shutdown successfully")
		}
	}()

	logger.Printf("🚀 Starting server at http://0.0.0.0:%s", cfg.Service.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer shutdownCtxCancel()

	isShutdownErrors := false

	if err = srv.Shutdown(shutdownCtx); err != nil {
		logger.Printf(err.Error())
		isShutdownErrors = true
	}

	if isShutdownErrors {
		logger.Printf("Server closed, but not all resources closed properly!")
	} else {
		logger.Printf("✅ Server shutdown successfully")
	}
}

package app

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/awakedx/task/internal/config"
	"github.com/awakedx/task/internal/controller"
	"github.com/awakedx/task/internal/controller/middleware"
	"github.com/awakedx/task/internal/repository"
	pg "github.com/awakedx/task/internal/repository"
	"github.com/awakedx/task/internal/service"
	"github.com/go-playground/validator/v10"
)

func StartServer() error {
	ctx := context.Background()

	cfg := config.Get()

	slog.Info("connection to DB")
	db, err := pg.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	validate := validator.New()
	store := repository.NewStore(db)
	services := service.NewService(store)
	handlers := controller.NewHandler(ctx, services, validate)

	srv := http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  cfg.ReadTimeOut,
		WriteTimeout: cfg.WriteTimeOut,
		Handler:      middleware.LoggingMW(handlers.RegisterRoutes()),
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	slog.Info("Server started")

	sigShutDown := make(chan os.Signal, 1)
	signal.Notify(sigShutDown, syscall.SIGTERM, syscall.SIGINT)
	<-sigShutDown
	slog.Info("gracefully shutting down")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("error occuring by shutting down server", "error", err)
	}
	db.Close()
	return nil
}

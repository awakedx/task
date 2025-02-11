package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/awakedx/task/internal/config"
	"github.com/awakedx/task/internal/controller"
	"github.com/awakedx/task/internal/controller/middleware"
	"github.com/awakedx/task/internal/repository"
	pg "github.com/awakedx/task/internal/repository"
	"github.com/awakedx/task/internal/service"
	"github.com/go-playground/validator/v10"
)

func StartServer() error {
	cfg := config.Get()

	slog.Info("connection to DB")
	db, err := pg.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	validate := validator.New(validator.WithRequiredStructEnabled())
	store := repository.NewStore(db)
	services := service.NewService(store)
	handlers := controller.NewHandler(services, validate)

	srv := http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  cfg.ReadTimeOut,
		WriteTimeout: cfg.WriteTimeOut,
		Handler:      middleware.LoggingMW(handlers.RegisterRoutes()),
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("Server started")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	sigShutDown := make(chan os.Signal, 1)
	signal.Notify(sigShutDown, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err := <-errCh:
		return err
	case <-sigShutDown:
		slog.Info("gracefully shutting down")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("error occured by shutting down server", "error", err)
		}
		return nil
	}
}

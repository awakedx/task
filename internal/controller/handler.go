package controller

import (
	"context"
	"net/http"

	"github.com/awakedx/task/internal/controller/middleware"
	"github.com/awakedx/task/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	ctx       context.Context
	services  *service.Service
	validator *validator.Validate
}

func NewHandler(ctx context.Context, service *service.Service, validator *validator.Validate) *Handler {
	return &Handler{
		ctx:       ctx,
		services:  service,
		validator: validator,
	}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /item", middleware.AdminMW(h.NewItem))
	mux.Handle("DELETE /item/{id}", middleware.AdminMW(h.DeleteItem))
	mux.Handle("GET /item/{id}", middleware.AdminMW(h.GetItem))
	mux.Handle("PATCH /item/{id}", middleware.AdminMW(h.UpdateItem))

	mux.Handle("POST /seller", middleware.AdminMW(h.NewSeller))
	return mux
}

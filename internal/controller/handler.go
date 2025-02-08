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
	service   *service.Service
	validator *validator.Validate
}

func NewHandler(ctx context.Context, service *service.Service, validator *validator.Validate) *Handler {
	return &Handler{
		ctx:       ctx,
		service:   service,
		validator: validator,
	}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /items", middleware.AdminMW(h.NewItem))
	mux.Handle("DELETE /items/{id}", middleware.AdminMW(h.DeleteItem))
	mux.Handle("GET /items/{id}", middleware.AdminMW(h.GetItem))
	mux.Handle("PATCH /items/{id}", middleware.AdminMW(h.UpdateItem))

	mux.Handle("POST /sellers", middleware.AdminMW(h.NewSeller))

	mux.Handle("POST /orders", middleware.AdminMW(h.NewOrder))
	return mux
}

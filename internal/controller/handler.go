package controller

import (
	"net/http"

	"github.com/awakedx/task/internal/controller/middleware"
	"github.com/awakedx/task/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   *service.Service
	validator *validator.Validate
}

func NewHandler(service *service.Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	//items
	mux.Handle("POST /items", middleware.AdminMW(h.NewItem))
	mux.Handle("GET /items/{id}", middleware.AdminMW(h.GetItem))
	mux.Handle("PATCH /items/{id}", middleware.AdminMW(h.UpdateItem))
	mux.Handle("DELETE /items/{id}", middleware.AdminMW(h.DeleteItem))

	//sellers
	mux.Handle("POST /sellers", middleware.AdminMW(h.NewSeller))
	mux.Handle("GET /sellers/{id}", middleware.AdminMW(h.GetSeller))
	mux.Handle("PATCH /sellers/{id}", middleware.AdminMW(h.UpdateSeller))
	mux.Handle("DELETE /sellers/{id}", middleware.AdminMW(h.DeleteSeller))

	//customers
	mux.Handle("POST /customers", middleware.AdminMW(h.NewCustomer))
	mux.Handle("GET /customers/{id}", middleware.AdminMW(h.GetCustomer))
	mux.Handle("PATCH /customers/{id}", middleware.AdminMW(h.UpdateCustomer))
	mux.Handle("DELETE /customers/{id}", middleware.AdminMW(h.DeleteCustomer))

	//orders
	mux.Handle("POST /orders", middleware.AdminMW(h.NewOrder))
	mux.Handle("GET /orders/{id}", middleware.AdminMW(h.GetOrder))
	//TODO
	//mux.Handle("DELETE /orders/{id}", middleware.AdminMW(h.DeleteOrder)))
	//mux.Handle("PATCH /orders/{id}", middleware.AdminMW(h.UpdateOrder)))
	return mux
}

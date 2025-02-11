package controller

import (
	"encoding/json"
	"net/http"

	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/utils"
)

func (h *Handler) NewCustomer(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	err = h.validator.Struct(&customer)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	id, err := h.service.Customers.NewCustomer(r.Context(), &customer)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]any{
		"message": "successfully created",
		"id":      id,
	})
	return
}

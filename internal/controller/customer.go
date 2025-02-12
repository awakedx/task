package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	common "github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/utils"
	"github.com/google/uuid"
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

func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	customer, err := h.service.Customers.GetCustomer(r.Context(), id)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"eror": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"customer": customer,
	})
	return
}
func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	var updateCustomer common.UpdateCustomer
	err = json.NewDecoder(r.Body).Decode(&updateCustomer)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	err = h.validator.Struct(&updateCustomer)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	updateCustomer.Id = id
	fmt.Println(updateCustomer)
	err = h.service.Customers.UpdateCustomer(r.Context(), &updateCustomer)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"message": "successfully updated",
	})
	return
}
func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	err = h.service.Customers.DeleteCustomer(r.Context(), id)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"message": "successfully deleted",
	})
	return
}

package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/awakedx/task/internal/service/order"
	"github.com/awakedx/task/internal/utils"
)

func (h *Handler) NewOrder(w http.ResponseWriter, r *http.Request) {
	var orderD order.OrderDetails
	if err := json.NewDecoder(r.Body).Decode(&orderD); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	if err := h.validator.Struct(orderD); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	id, err := h.service.Orders.NewOrder(r.Context(), &orderD)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, map[string]any{
		"id": id,
	})
	return
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}
	order, err := h.service.Orders.GetById(r.Context(), orderId)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"order": order,
	})
	return
}

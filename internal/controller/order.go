package controller

import (
	"encoding/json"
	"net/http"

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
}

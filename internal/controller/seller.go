package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/utils"
)

func (h *Handler) NewSeller(w http.ResponseWriter, r *http.Request) {
	var newSeller domain.Seller
	if err := json.NewDecoder(r.Body).Decode(&newSeller); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	if err := h.validator.Struct(&newSeller); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	id, err := h.service.Sellers.Create(r.Context(), &newSeller)

	if err != nil && errors.Is(err, utils.BadRequestErr) {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"err:": err.Error(),
		})
		return
	} else if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"err:": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]any{
		"Message": "successfully created",
		"id":      id,
	})
	return
}

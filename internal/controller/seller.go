package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	common "github.com/awakedx/task/internal/common/update"
	"github.com/awakedx/task/internal/domain"
	"github.com/awakedx/task/internal/utils"
	"github.com/google/uuid"
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
		"message": "successfully created",
		"id":      id,
	})
	return
}

func (h *Handler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"err": err.Error(),
		})
		return
	}
	err = h.service.Sellers.Delete(r.Context(), id)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"err": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"message": "successfully deleted",
	})
}

func (h *Handler) GetSeller(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"err": err.Error(),
		})
		return
	}

	seller, err := h.service.Sellers.Get(r.Context(), id)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"err": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"seller": seller,
	})
}

func (h *Handler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	var updateSeller common.UpdateSeller
	err := json.NewDecoder(r.Body).Decode(&updateSeller)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"err": err.Error(),
		})
		return
	}
	id, err := uuid.Parse(r.PathValue("id"))
	updateSeller.Id = &id
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"err": err.Error(),
		})
		return
	}
	err = h.validator.Struct(&updateSeller)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"err": err.Error(),
		})
		return
	}
	err = h.service.Sellers.Update(r.Context(), &updateSeller)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"err": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"message": "successfully updated",
	})
}

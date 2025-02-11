package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/awakedx/task/internal/common/update"

	"github.com/awakedx/task/internal/service/item"
	"github.com/awakedx/task/internal/utils"
)

func (h *Handler) NewItem(w http.ResponseWriter, r *http.Request) {
	var itemValue item.ItemValues
	if err := json.NewDecoder(r.Body).Decode(&itemValue); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	if err := h.validator.Struct(&itemValue); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	ids, err := h.service.Items.NewItem(r.Context(), &itemValue)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]any{
		"Message": "successfully created",
		"ids":     ids,
	})
	return
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": "Invalid id",
		})
	}
	id, err := h.service.Items.Delete(r.Context(), itemId)
	if err != nil && err.Error() != "Nothing to delete" {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"delete id": id,
			"error":     err.Error(),
		})
		return
	}
	if err != nil && err.Error() == "Nothing to delete" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": "No item for deletion",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"deleted id": id,
	})
	return
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	itemId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	i, err := h.service.Items.Get(r.Context(), itemId)
	if err != nil && errors.Is(err, utils.NotFoundError) {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]any{
			"error": err.Error(),
		})
		return
	}
	if err != nil && errors.Is(err, utils.InternalError) {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"Item": i,
	})
	return
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	var updateItem common.UpdateItem
	if err = json.NewDecoder(r.Body).Decode(&updateItem); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	if err = h.validator.Struct(&updateItem); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	updateItem.Id = &id
	err = h.service.Items.UpdateItem(r.Context(), &updateItem)
	if err != nil && err.Error() != "not found by id" {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	if err != nil && err.Error() == "not found by id" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]any{
		"message": "successfully updated",
	})
	return
}

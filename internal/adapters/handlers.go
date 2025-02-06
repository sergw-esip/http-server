package adapters

import (
	"encoding/json"
	"httpServer/internal/cases"
	"httpServer/internal/entities"
	"net/http"
	"strconv"
)

type HTTPHandler struct {
	UseCases *cases.EntityUseCases
}

func NewHTTPHandler(useCases *cases.EntityUseCases) *HTTPHandler {
	return &HTTPHandler{UseCases: useCases}
}

func (h *HTTPHandler) CreateEntity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var entity entities.Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.UseCases.CreateEntity(&entity); err != nil {
		http.Error(w, "failed to create entity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity)
}

func (h *HTTPHandler) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	entityID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.UseCases.DeleteEntity(entityID); err != nil {
		http.Error(w, "failed to delete entity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandler) ListEntities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	entities, err := h.UseCases.ListEntities()
	if err != nil {
		http.Error(w, "failed to list entities", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entities)
}

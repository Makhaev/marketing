package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Makhaev/marketing/internal/repository"
)

type StoreHandler struct {
	Repo *repository.StoreRepository
}

func NewStoreHandler(repo *repository.StoreRepository) *StoreHandler {
	return &StoreHandler{Repo: repo}
}

func (h *StoreHandler) GetAllStores(w http.ResponseWriter, r *http.Request) {
	stores, err := h.Repo.GetAllStores()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Если нет магазинов, возвращаем пустой массив вместо null
	if stores == nil {
		stores = []repository.Store{}
	}

	json.NewEncoder(w).Encode(stores)
}

func (h *StoreHandler) CreateStore(w http.ResponseWriter, r *http.Request) {
	var s repository.Store
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if err := h.Repo.CreateStore(&s); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(s)
}

func (h *StoreHandler) UpdateStore(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid store id", http.StatusBadRequest)
		return
	}

	var s repository.Store
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	s.ID = id
	if err := h.Repo.UpdateStore(&s); err != nil {
		http.Error(w, "Failed to update store", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func (h *StoreHandler) DeleteStore(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid store id", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteStore(id); err != nil {
		http.Error(w, "Failed to delete store", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

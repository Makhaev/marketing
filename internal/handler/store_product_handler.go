package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Makhaev/marketing/internal/repository"
	"github.com/lib/pq"
)

type StoreProductHandler struct {
	Repo *repository.StoreProductRepository
}

func NewStoreProductHandler(repo *repository.StoreProductRepository) *StoreProductHandler {
	return &StoreProductHandler{Repo: repo}
}

func (h *StoreProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var sp repository.StoreProduct
	if err := json.NewDecoder(r.Body).Decode(&sp); err != nil {
		http.Error(w, `{"error":"Invalid JSON body"}`, http.StatusBadRequest)
		return
	}

	err := h.Repo.Create(&sp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23503":
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Store or Product does not exist"})
				return
			case "23505":
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(map[string]string{"error": "This product already exists in the store"})
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sp)
}
func (h *StoreProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *StoreProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/store-products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var sp repository.StoreProduct
	if err := json.NewDecoder(r.Body).Decode(&sp); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	sp.ID = id

	if err := h.Repo.Update(&sp); err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sp)
}

func (h *StoreProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/store-products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Makhaev/marketing/internal/repository"
)

type ProductHandler struct {
	Repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repo.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p repository.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateProduct(&p); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request, id int) {
	p, err := h.Repo.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request, id int) {
	var p repository.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := h.Repo.UpdateProduct(&p); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.Repo.DeleteProduct(id); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) GetPrices(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	productID, _ := strconv.Atoi(parts[1])
	prices, err := h.Repo.GetPricesByProduct(productID)
	if err != nil {
		http.Error(w, "Failed to fetch prices", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}

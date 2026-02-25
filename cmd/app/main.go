package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Makhaev/marketing/internal/db"
	"github.com/Makhaev/marketing/internal/handler"
	"github.com/Makhaev/marketing/internal/repository"
	"github.com/Makhaev/marketing/migrations"
)

func main() {
	database := db.InitPostgres()
	migrations.RunMigrations(database)

	// ===== Репозитории =====
	productRepo := repository.NewProductRepository(database)
	productHandler := handler.NewProductHandler(productRepo)

	storeRepo := repository.NewStoreRepository(database)
	storeHandler := handler.NewStoreHandler(storeRepo)

	storeProductRepo := repository.NewStoreProductRepository(database)
	storeProductHandler := handler.NewStoreProductHandler(storeProductRepo)

	mux := http.NewServeMux()

	userRepo := repository.NewUserRepository(database)
	userHandler := handler.NewUserHandler(userRepo)

	// ===== API USERS =====
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	//        PRODUCTS
	// =========================
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetAllProducts(w, r)
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProduct(w, r, id)
		case http.MethodPut:
			productHandler.UpdateProduct(w, r, id)
		case http.MethodDelete:
			productHandler.DeleteProduct(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	//        STORES
	// =========================
	mux.HandleFunc("/api/stores", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			storeHandler.GetAllStores(w, r)
		case http.MethodPost:
			storeHandler.CreateStore(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/stores/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			storeHandler.UpdateStore(w, r)
		case http.MethodDelete:
			storeHandler.DeleteStore(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	//     STORE PRODUCTS
	// =========================
	mux.HandleFunc("/api/store-products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			storeProductHandler.GetAll(w, r)
		case http.MethodPost:
			storeProductHandler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/store-products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			storeProductHandler.Update(w, r)
		case http.MethodDelete:
			storeProductHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// =========================
	//        REACT SPA
	// =========================
	distDir := "./frontend/dist"
	fs := http.FileServer(http.Dir(distDir))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		path := distDir + r.URL.Path
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			fs.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, distDir+"/index.html")
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package repository

import (
	"database/sql"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) CreateProduct(p *Product) error {
	query := `INSERT INTO products (name, category) VALUES ($1, $2) RETURNING id`
	return r.DB.QueryRow(query, p.Name, p.Category).Scan(&p.ID)
}

func (r *ProductRepository) GetAllProducts() ([]Product, error) {
	rows, err := r.DB.Query(`SELECT id, name, category FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (*Product, error) {
	var p Product
	err := r.DB.QueryRow(`SELECT id, name, category FROM products WHERE id=$1`, id).Scan(&p.ID, &p.Name, &p.Category)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) UpdateProduct(p *Product) error {
	_, err := r.DB.Exec(`UPDATE products SET name=$1, category=$2 WHERE id=$3`, p.Name, p.Category, p.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	_, err := r.DB.Exec(`DELETE FROM products WHERE id=$1`, id)
	return err
}

// Для GetPrices
type StorePrice struct {
	Store    string  `json:"store"`
	Price    float64 `json:"price"`
	IsPromo  bool    `json:"is_promo"`
	ImageURL *string `json:"image_url"`
}

func (r *ProductRepository) GetPricesByProduct(productID int) ([]StorePrice, error) {
	query := `
	SELECT s.name, sp.price, sp.is_promo, sp.image_url
	FROM store_products sp
	JOIN stores s ON sp.store_id = s.id
	WHERE sp.product_id = $1
	ORDER BY sp.price ASC`
	rows, err := r.DB.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prices []StorePrice
	for rows.Next() {
		var sp StorePrice
		if err := rows.Scan(&sp.Store, &sp.Price, &sp.IsPromo, &sp.ImageURL); err != nil {
			return nil, err
		}
		prices = append(prices, sp)
	}
	return prices, nil
}

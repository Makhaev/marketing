package repository

import (
	"database/sql"
	"fmt"
)

type StoreProduct struct {
	ID        int     `json:"id"`
	StoreID   int     `json:"store_id"`
	ProductID int     `json:"product_id"`
	Price     float64 `json:"price"`
	IsPromo   bool    `json:"is_promo"`
	ImageURL  *string `json:"image_url"`
}

type StoreProductRepository struct {
	DB *sql.DB
}

func NewStoreProductRepository(db *sql.DB) *StoreProductRepository {
	return &StoreProductRepository{DB: db}
}
func (r *StoreProductRepository) Create(sp *StoreProduct) error {
	query := `
	INSERT INTO store_products (store_id, product_id, price, is_promo, image_url)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	err := r.DB.QueryRow(query,
		sp.StoreID,
		sp.ProductID,
		sp.Price,
		sp.IsPromo,
		sp.ImageURL,
	).Scan(&sp.ID)
	if err != nil {
		return fmt.Errorf("failed to insert store_product: %w", err)
	}
	return nil
}

func (r *StoreProductRepository) GetAll() ([]StoreProduct, error) {
	rows, err := r.DB.Query(`
		SELECT id, store_id, product_id, price, is_promo, image_url
		FROM store_products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []StoreProduct
	for rows.Next() {
		var sp StoreProduct
		if err := rows.Scan(
			&sp.ID,
			&sp.StoreID,
			&sp.ProductID,
			&sp.Price,
			&sp.IsPromo,
			&sp.ImageURL,
		); err != nil {
			return nil, err
		}
		list = append(list, sp)
	}
	return list, nil
}

func (r *StoreProductRepository) Update(sp *StoreProduct) error {
	_, err := r.DB.Exec(`
		UPDATE store_products
		SET store_id=$1, product_id=$2, price=$3, is_promo=$4, image_url=$5
		WHERE id=$6`,
		sp.StoreID,
		sp.ProductID,
		sp.Price,
		sp.IsPromo,
		sp.ImageURL,
		sp.ID,
	)
	return err
}

func (r *StoreProductRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM store_products WHERE id=$1`, id)
	return err
}

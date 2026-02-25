package repository

import "database/sql"

type Store struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	OwnerID int    `json:"owner_id"`
}
type StoreRepository struct {
	DB *sql.DB
}

func NewStoreRepository(db *sql.DB) *StoreRepository {
	return &StoreRepository{DB: db}
}
func (r *StoreRepository) CreateStore(s *Store) error {
	query := `INSERT INTO stores (name, address, owner_id) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, s.Name, s.Address, s.OwnerID).Scan(&s.ID)
}

func (r *StoreRepository) GetAllStores() ([]Store, error) {
	rows, err := r.DB.Query(`SELECT id, name, address, owner_id FROM stores`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []Store
	for rows.Next() {
		var s Store
		if err := rows.Scan(&s.ID, &s.Name, &s.Address, &s.OwnerID); err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}
	return stores, nil
}

func (r *StoreRepository) UpdateStore(s *Store) error {
	_, err := r.DB.Exec(
		`UPDATE stores SET name=$1, address=$2 WHERE id=$3`,
		s.Name, s.Address, s.ID,
	)
	return err
}

func (r *StoreRepository) DeleteStore(id int) error {
	_, err := r.DB.Exec(`DELETE FROM stores WHERE id=$1`, id)
	return err
}

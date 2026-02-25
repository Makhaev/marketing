package repository

import (
	"database/sql"
)

type User struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(phone, role string) (*User, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO users (phone, role) VALUES ($1, $2) RETURNING id",
		phone, role,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Phone: phone, Role: role}, nil
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT id, phone, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Phone, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

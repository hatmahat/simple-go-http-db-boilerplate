package repository

import (
	"database/sql"
	"simple-http-boilerplate/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (r *UserRepo) Create(user model.User) (*model.User, error) {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	if err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID); err != nil {
		return nil, err
	}
	return &user, nil
}

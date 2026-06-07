package repository

import (
    "database/sql"
    "expense-tracker/internal/models"
)

type UserRepo struct{ db *sql.DB }
func NewUserRepo(db *sql.DB) *UserRepo { return &UserRepo{db} }

func (r *UserRepo) Create(name, email, hash string) (*models.User, error) {
    var u models.User
    err := r.db.QueryRow(
        `INSERT INTO users (name, email, password) VALUES ($1,$2,$3)
         RETURNING id, name, email, created_at`,
        name, email, hash,
    ).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
    return &u, err
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
    var u models.User
    err := r.db.QueryRow(
        `SELECT id, name, email, password FROM users WHERE email=$1`, email,
    ).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
    return &u, err
}
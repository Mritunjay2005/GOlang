package db

import (
    "database/sql"
    "fmt"

    "expense-tracker/internal/config"
    _ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )
    return sql.Open("postgres", dsn)
}
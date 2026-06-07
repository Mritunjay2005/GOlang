package db

import (
    "database/sql"
    "log"
)

func Migrate(db *sql.DB) {
    stmts := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id         SERIAL PRIMARY KEY,
            name       VARCHAR(100) NOT NULL,
            email      VARCHAR(150) UNIQUE NOT NULL,
            password   TEXT         NOT NULL,
            created_at TIMESTAMPTZ  DEFAULT NOW()
        )`,
        `CREATE TABLE IF NOT EXISTS expenses (
            id           SERIAL PRIMARY KEY,
            user_id      INT           NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            title        VARCHAR(200)  NOT NULL,
            amount       NUMERIC(12,2) NOT NULL CHECK (amount > 0),
            category     VARCHAR(50)   NOT NULL,
            note         TEXT,
            expense_date DATE          NOT NULL DEFAULT CURRENT_DATE,
            created_at   TIMESTAMPTZ   DEFAULT NOW(),
            updated_at   TIMESTAMPTZ   DEFAULT NOW()
        )`,
        `CREATE INDEX IF NOT EXISTS idx_expenses_user_id ON expenses(user_id)`,
    }
    for _, s := range stmts {
        if _, err := db.Exec(s); err != nil {
            log.Fatalf("migration failed: %v", err)
        }
    }
    log.Println("Migrations applied successfully.")
}
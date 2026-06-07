package models

import "time"

var Categories = []string{
    "Food", "Transport", "Health", "Shopping",
    "Housing", "Entertainment", "Education", "Other",
}

type Expense struct {
    ID          int       `json:"id"`
    UserID      int       `json:"user_id"`
    Title       string    `json:"title"`
    Amount      float64   `json:"amount"`
    Category    string    `json:"category"`
    Note        string    `json:"note"`
    ExpenseDate string    `json:"expense_date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
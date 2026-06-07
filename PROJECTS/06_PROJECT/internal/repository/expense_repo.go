package repository

import (
    "database/sql"
    "expense-tracker/internal/models"
    "fmt"
)

type ExpenseRepo struct{ db *sql.DB }
func NewExpenseRepo(db *sql.DB) *ExpenseRepo { return &ExpenseRepo{db} }

func (r *ExpenseRepo) List(userID int, category, from, to string) ([]models.Expense, error) {
    query := `SELECT id, user_id, title, amount, category, COALESCE(note,''),
                     expense_date::text, created_at, updated_at
              FROM expenses WHERE user_id=$1`
    args := []interface{}{userID}
    i := 2
    if category != "" {
        query += fmt.Sprintf(" AND category=$%d", i); args = append(args, category); i++
    }
    if from != "" {
        query += fmt.Sprintf(" AND expense_date>=$%d", i); args = append(args, from); i++
    }
    if to != "" {
        query += fmt.Sprintf(" AND expense_date<=$%d", i); args = append(args, to); i++
    }
    query += " ORDER BY expense_date DESC"
    rows, err := r.db.Query(query, args...)
    if err != nil { return nil, err }
    defer rows.Close()
    var expenses []models.Expense
    for rows.Next() {
        var e models.Expense
        rows.Scan(&e.ID, &e.UserID, &e.Title, &e.Amount, &e.Category,
                  &e.Note, &e.ExpenseDate, &e.CreatedAt, &e.UpdatedAt)
        expenses = append(expenses, e)
    }
    return expenses, nil
}

func (r *ExpenseRepo) Create(e models.Expense) (*models.Expense, error) {
    err := r.db.QueryRow(
        `INSERT INTO expenses (user_id,title,amount,category,note,expense_date)
         VALUES ($1,$2,$3,$4,$5,$6)
         RETURNING id, created_at, updated_at`,
        e.UserID, e.Title, e.Amount, e.Category, e.Note, e.ExpenseDate,
    ).Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt)
    return &e, err
}

func (r *ExpenseRepo) Update(e models.Expense) (*models.Expense, error) {
    err := r.db.QueryRow(
        `UPDATE expenses SET title=$1,amount=$2,category=$3,note=$4,
         expense_date=$5,updated_at=NOW()
         WHERE id=$6 AND user_id=$7
         RETURNING updated_at`,
        e.Title, e.Amount, e.Category, e.Note, e.ExpenseDate, e.ID, e.UserID,
    ).Scan(&e.UpdatedAt)
    return &e, err
}

func (r *ExpenseRepo) Delete(id, userID int) error {
    _, err := r.db.Exec(`DELETE FROM expenses WHERE id=$1 AND user_id=$2`, id, userID)
    return err
}

func (r *ExpenseRepo) Summary(userID int) (map[string]interface{}, error) {
    var total float64
    r.db.QueryRow(`SELECT COALESCE(SUM(amount),0) FROM expenses WHERE user_id=$1`, userID).Scan(&total)
    rows, err := r.db.Query(
        `SELECT category, SUM(amount) FROM expenses WHERE user_id=$1 GROUP BY category`, userID)
    if err != nil { return nil, err }
    defer rows.Close()
    byCategory := map[string]float64{}
    for rows.Next() {
        var cat string; var amt float64
        rows.Scan(&cat, &amt)
        byCategory[cat] = amt
    }
    return map[string]interface{}{"total": total, "by_category": byCategory}, nil
}
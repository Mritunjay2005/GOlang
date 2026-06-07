package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "expense-tracker/internal/middleware"
    "expense-tracker/internal/models"
    "expense-tracker/internal/repository"
)

func Expenses(db *sql.DB) http.HandlerFunc {
    repo := repository.NewExpenseRepo(db)
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value(middleware.UserIDKey).(int)
        w.Header().Set("Content-Type", "application/json")

        switch r.Method {
        case http.MethodGet:
            cat  := r.URL.Query().Get("category")
            from := r.URL.Query().Get("from")
            to   := r.URL.Query().Get("to")
            expenses, err := repo.List(userID, cat, from, to)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(expenses)

        case http.MethodPost:
            var e models.Expense
            if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
                http.Error(w, "bad request", http.StatusBadRequest)
                return
            }
            e.UserID = userID
            created, err := repo.Create(e)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(created)
        }
    }
}

func ExpenseByID(db *sql.DB) http.HandlerFunc {
    repo := repository.NewExpenseRepo(db)
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value(middleware.UserIDKey).(int)
        parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/expenses/"), "/")
        if len(parts) == 0 || parts[0] == "" || parts[0] == "summary" {
            Summary(db)(w, r)
            return
        }
        id, err := strconv.Atoi(parts[0])
        if err != nil {
            http.Error(w, "invalid id", http.StatusBadRequest)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        switch r.Method {
        case http.MethodPut:
            var e models.Expense
            json.NewDecoder(r.Body).Decode(&e)
            e.ID     = id
            e.UserID = userID
            updated, err := repo.Update(e)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(updated)
        case http.MethodDelete:
            if err := repo.Delete(id, userID); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
        }
    }
}

func Summary(db *sql.DB) http.HandlerFunc {
    repo := repository.NewExpenseRepo(db)
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value(middleware.UserIDKey).(int)
        summary, err := repo.Summary(userID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(summary)
    }
}
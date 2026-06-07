package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "time"

    "expense-tracker/internal/config"
    "expense-tracker/internal/repository"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }
        var req struct {
            Name     string `json:"name"`
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        repo := repository.NewUserRepo(db)
        user, err := repo.Create(req.Name, req.Email, string(hash))
        if err != nil {
            http.Error(w, "email already exists", http.StatusConflict)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{"id": user.ID, "name": user.Name})
    }
}

func Login(db *sql.DB, cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        repo := repository.NewUserRepo(db)
        user, err := repo.FindByEmail(req.Email)
        if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
            http.Error(w, "invalid credentials", http.StatusUnauthorized)
            return
        }
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user_id": user.ID,
            "exp":     time.Now().Add(24 * time.Hour).Unix(),
        })
        tokenStr, _ := token.SignedString([]byte(cfg.JWTSecret))
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
    }
}
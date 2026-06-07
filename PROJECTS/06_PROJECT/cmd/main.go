package main

import (
    "log"
    "net/http"

    "expense-tracker/internal/config"
    "expense-tracker/internal/db"
    "expense-tracker/internal/handlers"
    "expense-tracker/internal/middleware"

    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()

    cfg := config.Load()

    database, err := db.Connect(cfg)
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }
    defer database.Close()

    db.Migrate(database)

    mux := http.NewServeMux()

    // Static frontend
    fs := http.FileServer(http.Dir(cfg.FrontendDir))
    mux.Handle("/", fs)

    // Auth (public)
    mux.HandleFunc("/api/register", handlers.Register(database))
    mux.HandleFunc("/api/login",    handlers.Login(database, cfg))

    // Expenses (protected)
    mux.Handle("/api/expenses",
        middleware.Auth(cfg)(http.HandlerFunc(handlers.Expenses(database))))
    mux.Handle("/api/expenses/",
        middleware.Auth(cfg)(http.HandlerFunc(handlers.ExpenseByID(database))))
    mux.Handle("/api/expenses/summary",
        middleware.Auth(cfg)(http.HandlerFunc(handlers.Summary(database))))

    log.Printf("Server listening on :%s", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, mux))
}
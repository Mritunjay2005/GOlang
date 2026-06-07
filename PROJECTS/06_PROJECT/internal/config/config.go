package config

import "os"

type Config struct {
    DBHost      string
    DBPort      string
    DBUser      string
    DBPassword  string
    DBName      string
    JWTSecret   string
    ServerPort  string
    FrontendDir string
}

func Load() *Config {
    return &Config{
        DBHost:      getEnv("DB_HOST", "localhost"),
        DBPort:      getEnv("DB_PORT", "5432"),
        DBUser:      getEnv("DB_USER", "expenseuser"),
        DBPassword:  getEnv("DB_PASSWORD", "expensepass"),
        DBName:      getEnv("DB_NAME", "expensedb"),
        JWTSecret:   getEnv("JWT_SECRET", "secret"),
        ServerPort:  getEnv("SERVER_PORT", "8080"),
        FrontendDir: getEnv("FRONTEND_DIR", "./frontend"),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
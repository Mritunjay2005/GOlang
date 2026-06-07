# 💰 Expense Tracker

A fully functional **full-stack expense tracking web application** built with Go, PostgreSQL, and Vanilla JavaScript. Features JWT authentication, complete CRUD operations, category filtering, and real-time spending charts.

---

## 🚀 Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.22+ (net/http stdlib) |
| Database | PostgreSQL 16 |
| Authentication | JWT + bcrypt |
| Frontend | HTML5 / CSS3 / Vanilla JavaScript |
| Charts | Chart.js 4.4 |
| Config | godotenv |

---

## ✨ Features

- 🔐 **User Authentication** — Register & login with bcrypt-hashed passwords and 24-hour JWT tokens
- 📝 **Expense Management** — Add, edit, delete expenses with title, amount, category, date & notes
- 🔍 **Smart Filtering** — Filter by category and custom date ranges
- 📊 **Dashboard Summary** — Total spent, monthly total, top spending category
- 📈 **Charts** — Interactive pie & bar charts showing spending breakdown by category
- 🗄️ **Auto Migrations** — Database tables created automatically on first server start

---

## 📁 Project Structure

```
expense-tracker/
├── cmd/
│   └── main.go                  # Entry point — starts HTTP server
├── internal/
│   ├── config/
│   │   └── config.go            # Loads env vars into Config struct
│   ├── db/
│   │   ├── db.go                # Opens PostgreSQL connection pool
│   │   └── migrations.go        # Runs CREATE TABLE SQL on startup
│   ├── models/
│   │   ├── user.go              # User struct
│   │   └── expense.go           # Expense struct + Category list
│   ├── repository/
│   │   ├── user_repo.go         # DB queries for users
│   │   └── expense_repo.go      # DB queries for expenses (CRUD)
│   ├── handlers/
│   │   ├── auth_handler.go      # POST /api/register, POST /api/login
│   │   └── expense_handler.go   # CRUD endpoints for expenses
│   └── middleware/
│       └── auth.go              # JWT validation middleware
├── frontend/
│   ├── index.html               # Login / Register page
│   ├── dashboard.html           # Main app page
│   ├── css/
│   │   └── style.css            # All styles
│   └── js/
│       ├── auth.js              # Login / register API calls
│       ├── expenses.js          # Expense CRUD + table rendering
│       └── charts.js            # Chart.js pie/bar charts
├── .env                         # Environment variables
├── .env.example                 # Template for .env
└── go.mod                       # Go module definition
```

---

## ⚙️ Prerequisites

- [Go 1.22+](https://golang.org/dl/)
- [PostgreSQL 16+](https://www.postgresql.org/download/)
- Git Bash or any terminal

---

## 🛠️ Setup & Installation

### 1. Clone the repository

```bash
git clone https://github.com/your-username/expense-tracker.git
cd expense-tracker
```

### 2. Set up PostgreSQL

```sql
CREATE USER expenseuser WITH PASSWORD 'expensepass';
CREATE DATABASE expensedb OWNER expenseuser;
GRANT ALL PRIVILEGES ON DATABASE expensedb TO expenseuser;
```

### 3. Configure environment variables

```bash
cp .env.example .env
```

Edit `.env` with your values:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=expenseuser
DB_PASSWORD=expensepass
DB_NAME=expensedb
JWT_SECRET=your_long_random_secret_here
SERVER_PORT=8080
FRONTEND_DIR=./frontend
```

### 4. Install Go dependencies

```bash
go mod tidy
```

### 5. Run the server

```bash
go run ./cmd/main.go
```

### 6. Open in browser

```
http://localhost:8080
```

---

## 🗃️ Database Schema

```sql
-- Users table
CREATE TABLE users (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(150) UNIQUE NOT NULL,
    password   TEXT         NOT NULL,   -- bcrypt hash
    created_at TIMESTAMPTZ  DEFAULT NOW()
);

-- Expenses table
CREATE TABLE expenses (
    id           SERIAL PRIMARY KEY,
    user_id      INT           NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title        VARCHAR(200)  NOT NULL,
    amount       NUMERIC(12,2) NOT NULL CHECK (amount > 0),
    category     VARCHAR(50)   NOT NULL,
    note         TEXT,
    expense_date DATE          NOT NULL DEFAULT CURRENT_DATE,
    created_at   TIMESTAMPTZ   DEFAULT NOW(),
    updated_at   TIMESTAMPTZ   DEFAULT NOW()
);
```

---

## 🔌 API Endpoints

### Auth (Public)

| Method | Endpoint | Body |
|--------|----------|------|
| `POST` | `/api/register` | `{ "name", "email", "password" }` |
| `POST` | `/api/login` | `{ "email", "password" }` |

### Expenses (JWT Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/expenses` | List all expenses (supports `?category=&from=&to=`) |
| `POST` | `/api/expenses` | Create a new expense |
| `PUT` | `/api/expenses/{id}` | Update an expense |
| `DELETE` | `/api/expenses/{id}` | Delete an expense |
| `GET` | `/api/expenses/summary` | Get total & breakdown by category |

### Example Request

```bash
# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret"}'

# Add expense (with token)
curl -X POST http://localhost:8080/api/expenses \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Lunch","amount":150.00,"category":"Food","expense_date":"2024-06-01"}'
```

---

## 🗂️ Expense Categories

`Food` · `Transport` · `Health` · `Shopping` · `Housing` · `Entertainment` · `Education` · `Other`

---

## 🏗️ Build for Production

```bash
go build -o bin/expense-tracker ./cmd/main.go
./bin/expense-tracker
```

---

## 🔮 Future Improvements

- [ ] Export expenses to CSV
- [ ] Monthly budget limits per category
- [ ] Dark mode toggle
- [ ] Recurring expense support
- [ ] Email alerts when budget exceeded
- [ ] Mobile-responsive improvements

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

> Built as a **Minor Project** demonstrating full-stack development with Go, PostgreSQL, and Vanilla JavaScript.

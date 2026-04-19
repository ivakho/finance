# Finance Tracker Backend (Go)

This is a backend service for a personal finance management application.

The system allows you to:
- Manage **categories**
- Track **income and expenses**
- View totals and transactions
- Interact with the system via **REST API** or a **Telegram Bot**

The entire project is written in **Golang** and uses:
- Gin (HTTP framework)
- PostgreSQL (database)
- Docker (for DB)
- Telegram Bot API (for chat interaction)

---

# 🚀 Features

- Categories CRUD
- Transactions CRUD
- Income / Expense separation
- Aggregated totals
- Telegram bot with full interaction flow
- Clean architecture (handler → usecase → repository → storage)

---

# ⚙️ Requirements

- Go 1.20+
- Docker + Docker Compose
- `migrate` CLI (https://github.com/golang-migrate/migrate)

---

## 🌐 Live Demo

The frontend application is available on GitHub Pages:

👉 https://ivakho.github.io/fincalc/

# 📦 Installation & Setup

## 1. Clone repository

```bash
git clone <repo_url>
cd finance 
```

### 2. Telegram Bot Setup

Create a bot via @BotFather in telegram (https://web.telegram.org/)
Get your telegram token

### 3. Configure .env

Create .env file:

```env
DB_HOST=localhost
DB_PORT=5442
DB_USERNAME=[your db username]
DB_PASSWORD=[your db pass]
DB_NAME=[your db name]
DB_SSLMODE=disable
APP_PORT=8080
TG_TOKEN=[your telegram token]
API_URL=http://localhost:8080
CORS_ALLOWED_ORIGINS=[http://localhost:5173, http://localhost:5174, etc.]
```

### 4. Start PostgreSQL (Docker)

```bash
docker compose up
```

Don't forget to make database migrations with first init

```bash
make migrate
```

### 5. Run backend server

```bash
go run cmd/main.go
```

Server will be available at:
http://localhost:8090

Bot Features

Main menu contains:

📉 Expenses
📈 Income
📂 Categories

You can:

Add transactions
Update transactions
Delete transactions
View lists
Manage categories

All actions are handled step-by-step via bot or frontend version which available at .

---

## 📡 API

### Categories
- POST /category/
- GET /category/
- GET /category/:id
- PUT /category/
- DELETE /category/:id

### Transactions
- POST /transactions/
- GET /transactions/:id
- GET /transactions/getAll
- GET /transactions/getIncome
- GET /transactions/getExpense
- PUT /transactions/
- DELETE /transactions/:id

---

## 🧪 Examples

### Create category
```bash
curl -X POST http://localhost:8090/category/ \
-H "Content-Type: application/json" \
-d '{"name": "Food"}'
```

### Create transaction
```bash
curl -X POST http://localhost:8090/transactions/ \
-H "Content-Type: application/json" \
-d '{"category_id":1,"type":"expense","amount":1000,"created_at":"2026-04-18"}'
```

---

## 🧱 Architecture

```
handler → usecase → repository → storage
```

---

## 📂 Structure

```
internal/
  api/
  usecase/
  repository/
  storage/
  service/

db/
  migrations/
```

---

## 🛠 Commands

```bash
make migrate
make tidy
```

---
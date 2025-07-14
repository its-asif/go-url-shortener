

# 🔗 URL Shortener

A simple URL shortener built with **Go**, **PostgreSQL**, **SQLX**, and **Gorilla Mux**.

---

## 🚀 Features

- Shorten any long URL
- Redirect to original URL using short code

---

## 🛠️ Setup

### 1. Clone the repo

```bash
git clone https://github.com/its-asif/go-url-shortener.git
cd go-url-shortener
```

### 2. Create `.env` file

```env
# .env
DB_URL=postgres://<username>:<password>@localhost:5432/<database>?sslmode=disable
```

You can refer to `.env.example` for the format.

### 3. Create the database table

```sql
CREATE TABLE urls (
  id SERIAL PRIMARY KEY,
  short_code VARCHAR(10) UNIQUE NOT NULL,
  original_url TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Run the project

```bash
go run main.go
```

---

## 📦 API Endpoints

### ➕ Shorten a URL

```http
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com"
}
```

### 🔁 Redirect

```http
GET /<short_code>
```

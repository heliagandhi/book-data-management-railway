Hi, saya Helia! <img src="https://media.giphy.com/media/mGcNjsfWAjY5AEZNw6/giphy.gif" width="50">

---

# 📚 Book Data Management API

API sederhana untuk mengelola data **Users**, **Categories**, dan **Books** menggunakan **Golang (Gin Framework)**, **PostgreSQL**, **JWT Authentication**, dan **Swagger Documentation**.

---

## 🚀 Features

- 🔐 **Authentication** (Login + JWT)
- 👤 **User Management** (CRUD)
- 📂 **Category Management** (CRUD)
- 📖 **Book Management** (CRUD)
- 🔎 **Get Books by Category**
- 📑 **Swagger Documentation**

---

## 🧱 Tech Stack

- Golang
- Gin Framework
- PostgreSQL
- JWT Authentication
- Swaggo / Swagger Docs

---

## 📁 Struktur Project

```
book-data-management-railway/
├── controllers/
├── models/
├── repositories/
├── services/
├── middleware/
├── config/
├── docs/        # Swagger docs
├── main.go
├── go.mod
```

---

## ⚙️ Installation & Setup

### 1. Clone Repository

```bash
git clone https://github.com/heliagandhi/book-data-management-railway.git
cd book-data-management-railway
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Database (PostgreSQL)

```sql
CREATE DATABASE book_data_management;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    release_year INT,
    price DECIMAL(12,2),
    total_page INT,
    thickness VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50),
    CONSTRAINT fk_category FOREIGN KEY(category_id)
    REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INT,
    token TEXT,
    expired_at TIMESTAMP
);
```

#### Edit file `config/config.go` sesuai environment:

```go
connStr := "host=localhost port=5432 user=postgres password=YOUR_PASSWORD dbname=book_data_management sslmode=disable"
```

### 4. Run Application

```bash
go run main.go
```

Jika berhasil:

```
Database connected
[GIN-debug] Listening and serving HTTP on :8080
```

Server berjalan di:

```
http://localhost:8080
```

---

## 🔐 Authentication

### Login

```
POST /api/users/login
```

**Request:**

```json
{
  "username": "admin",
  "password": "admin"
}
```

**Response:**

```json
{
  "accessToken": "JWT_TOKEN",
  "expiredAt": "2026-04-05T12:58:37+07:00",
  "message": "success",
  "status": "200"
}
```

Gunakan token di header:

```
Authorization: Bearer JWT_TOKEN
```
<img width="819" height="392" alt="image" src="https://github.com/user-attachments/assets/3297a4e5-9423-4d76-9572-076274c369df" />



---

## ⚠️ Validasi Input

- `username`, `password` wajib diisi saat register
- `title`, `category_id` wajib diisi saat membuat buku
- `release_year` harus antara 1980–2024
- `total_page` harus > 0
- `category_id` harus valid dan ada di database
- `thickness` buku otomatis:
  - '>' 100 halaman → `"tebal"`
  - '≤' 100 halaman → `"tipis"`

---

## 📌 API Endpoints

### 👤 User

| Method | Path                    | Protected | Deskripsi                       |
| ------ | ----------------------- | --------- | ------------------------------- |
| POST   | /api/users              | ❌        | Register user baru              |
| POST   | /api/users/login        | ❌        | Login user & dapatkan token JWT |
| GET    | /api/users              | ❌        | Ambil semua user                |
| GET    | /api/users/:id          | ❌        | Ambil detail user               |
| PUT    | /api/users/:id          | ❌        | Update username user            |
| PUT    | /api/users/:id/password | ❌        | Update password user            |
| DELETE | /api/users/:id          | ❌        | Hapus user                      |

---

### 📂 Category

| Method | Path                      | Protected | Deskripsi                                   |
| ------ | ------------------------- | --------- | ------------------------------------------- |
| GET    | /api/categories           | ✅        | Ambil semua kategori                        |
| POST   | /api/categories           | ✅        | Buat kategori baru                          |
| GET    | /api/categories/:id       | ✅        | Ambil detail kategori                       |
| PUT    | /api/categories/:id       | ✅        | Update nama kategori                        |
| DELETE | /api/categories/:id       | ✅        | Hapus kategori (buku terkait ikut terhapus) |
| GET    | /api/categories/:id/books | ✅        | Ambil semua buku berdasarkan kategori       |

---

### 📖 Book

| Method | Path           | Protected | Deskripsi         |
| ------ | -------------- | --------- | ----------------- |
| GET    | /api/books     | ✅        | Ambil semua buku  |
| POST   | /api/books     | ✅        | Buat buku baru    |
| GET    | /api/books/:id | ✅        | Ambil detail buku |
| PUT    | /api/books/:id | ✅        | Update buku       |
| DELETE | /api/books/:id | ✅        | Hapus buku        |

---

## 📝 Swagger Docs

```
http://localhost:8080/swagger/index.html
```

---

## ⚠️ Catatan

- Semua route yang **Protected** harus menggunakan JWT di header `Authorization: Bearer TOKEN`.
- Menghapus kategori → otomatis menghapus semua buku terkait (_cascade delete_).
- Field `thickness` di buku terisi otomatis berdasarkan `total_page`.
- Gunakan Swagger untuk mencoba endpoint langsung tanpa perlu curl atau Postman.

---

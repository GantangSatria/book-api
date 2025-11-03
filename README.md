## Book API – Go + Gin + PostgreSQL + JWT

Aplikasi REST API untuk mengelola **data buku dan kategori**, lengkap dengan **autentikasi JWT** dan arsitektur **Clean Architecture**.



### Struktur Folder

```
book-api/
├── cmd/
│   └── main.go
├── config/
│   └── db.go
├── internal/
│   ├── bootstrap/
│   ├── controllers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── route/
│   └── services/
├── migrations/
│   └── schema.sql
├── go.mod
├── go.sum
└── README.md
```

Cara Menjalankan Project

#### Buat Database PostgreSQL

Buat database di PostgreSQL:

```sql
CREATE DATABASE book_db;
```

Lalu jalankan migrasi:

```bash
psql -U postgres -d book_db -f migrations/schema.sql
```

---

#### Setup File `.env`

Buat file `.env` di root project dan isi:

```env
DATABASE_URL=postgres://postgres:password@localhost:5432/book_db?sslmode=disable
JWT_SECRET=yourkey
PORT=8080
```

---

#### Jalankan Server

```bash
go run cmd/main.go
```

Server akan berjalan di:

```
http://localhost:8080
```

---

### Autentikasi

Project ini menggunakan **JWT (JSON Web Token)**.

#### Register User

**POST** `/api/users/register`

**Request:**

```json
{
  "username": "satria",
  "password": "123456"
}
```

**Response:**

```json
{
  "message": "user created successfully",
  "user_id": 1
}
```

#### 📌 Login User

**POST** `/api/users/login`

**Request:**

```json
{
  "username": "satria",
  "password": "123456"
}
```

**Response:**

```json
{
  "token": "<JWT_TOKEN>"
}
```

Gunakan token di header untuk semua endpoint lain:

```
Authorization: Bearer <JWT_TOKEN>
```

---

### Endpoint Kategori (`/api/categories`)

| Method   | Endpoint                    | Deskripsi                                   |
| -------- | --------------------------- | ------------------------------------------- |
| `GET`    | `/api/categories`           | Menampilkan semua kategori                  |
| `POST`   | `/api/categories`           | Menambah kategori baru                      |
| `GET`    | `/api/categories/:id`       | Menampilkan detail kategori                 |
| `DELETE` | `/api/categories/:id`       | Menghapus kategori                          |
| `GET`    | `/api/categories/:id/books` | Menampilkan semua buku berdasarkan kategori |

**Contoh Request (POST):**

```json
{
  "name": "Fiksi"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Fiksi"
}
```

---

### Endpoint Buku (`/api/books`)

| Method   | Endpoint         | Deskripsi               |
| -------- | ---------------- | ----------------------- |
| `GET`    | `/api/books`     | Menampilkan semua buku  |
| `POST`   | `/api/books`     | Menambahkan buku baru   |
| `GET`    | `/api/books/:id` | Menampilkan detail buku |
| `DELETE` | `/api/books/:id` | Menghapus buku          |

**Validasi Buku:**

* `release_year` hanya boleh antara **1980 - 2024**
* `thickness` otomatis diisi:

  * `> 100 halaman → "tebal"`
  * `<= 100 halaman → "tipis"`

**Contoh Request (POST):**

```json
{
  "title": "Laskar Pelangi",
  "description": "Kisah inspiratif dari Belitung",
  "image_url": "https://example.com/laskar-pelangi.jpg",
  "release_year": 2008,
  "price": 80000,
  "total_page": 250,
  "category_id": 1
}
```

**Response:**

```json
{
  "id": 1,
  "title": "Laskar Pelangi",
  "thickness": "tebal",
  "category_id": 1
}
```

---

### Fitur Validasi & Error Handling

* JWT middleware pada semua endpoint `/api/books` dan `/api/categories`.
* Validasi input untuk kategori & buku.
* Pesan error user-friendly:

  * Jika `category_id` tidak ditemukan.
  * Jika `release_year` di luar batas (1980–2024).
  * Jika data dihapus tapi ID tidak ada.

---

### Tools & Dependencies

* [Gin Gonic](https://github.com/gin-gonic/gin)
* [PostgreSQL](https://www.postgresql.org/)
* [lib/pq](https://github.com/lib/pq)
* [golang-jwt/jwt/v4](https://github.com/golang-jwt/jwt)
* [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

### Author

**Gantang Satria Yudha**

---

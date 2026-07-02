# 📚 Golang-Web-API

REST API sederhana untuk manajemen data buku, dibangun dengan **Go**, **Gin**, dan **GORM**. Proyek ini menerapkan arsitektur berlapis (*layered architecture*) — handler, service, dan repository — untuk memisahkan tanggung jawab tiap komponen.

## ✨ Fitur

- **CRUD Buku** — Create, Read, Update, Delete data buku
- **Validasi Request** — Validasi input menggunakan `go-playground/validator`
- **Layered Architecture** — Pemisahan handler, service, repository, dan entity
- **ORM** — Menggunakan GORM untuk komunikasi dengan database MySQL

## 🛠️ Tech Stack

| Komponen   | Teknologi                |
|------------|---------------------------|
| Bahasa     | Go 1.23                   |
| Framework  | [Gin](https://github.com/gin-gonic/gin) |
| ORM        | [GORM](https://gorm.io/)  |
| Database   | MySQL                     |
| Validasi   | go-playground/validator   |

## 📁 Struktur Proyek

```
Golang-Web-API/
├── book/
│   ├── entity.go       # Struct Book (model database)
│   ├── request.go      # Struct request body + validasi
│   ├── response.go     # Struct response API
│   ├── repository.go   # Layer akses database
│   └── service.go      # Layer business logic
├── handler/
│   └── book.go         # Layer HTTP handler (controller)
├── main.go              # Entry point & routing
├── go.mod
└── go.sum
```

## 🔌 API Endpoints

Base URL: `/api/v1`

| Method | Endpoint      | Deskripsi              |
|--------|---------------|-------------------------|
| GET    | `/books`      | Ambil semua data buku   |
| GET    | `/books/:id`  | Ambil buku berdasarkan ID |
| POST   | `/books`      | Tambah buku baru        |
| PUT    | `/books/:id`  | Update data buku        |
| DELETE | `/books/:id`  | Hapus data buku         |

### Contoh Request Body (POST/PUT)

```json
{
  "title": "Cinta Monyet",
  "description": "Novel remaja tentang cinta pertama",
  "price": 50000,
  "rating": 9,
  "discount": 10
}
```

## 🚀 Instalasi & Menjalankan

### Prasyarat
- Go >= 1.23
- MySQL

### Langkah-langkah

1. **Clone repository**
   ```bash
   git clone https://github.com/fairuzulum/Golang-Web-API.git
   cd Golang-Web-API
   ```

2. **Buat database**
   ```sql
   CREATE DATABASE golang_web_api;
   ```
   Tabel akan dibuat otomatis melalui `AutoMigrate` saat aplikasi pertama kali dijalankan.

3. **Sesuaikan koneksi database**
   
   Edit DSN di `main.go` sesuai kredensial MySQL kamu:
   ```go
   dsn := "root:@tcp(127.0.0.1:3306)/golang_web_api?charset=utf8mb4&parseTime=True&loc=Local"
   ```

4. **Install dependencies**
   ```bash
   go mod tidy
   ```

5. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```
   Server berjalan di `http://localhost:8080`

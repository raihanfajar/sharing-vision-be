# Sharing Vision API

REST API untuk use case **Post Article** menggunakan Go + Gin + MySQL.

---

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin (`github.com/gin-gonic/gin`)
- **Database**: MySQL 8.0
- **Driver**: `github.com/go-sql-driver/mysql`
- **Query style**: Raw SQL (no ORM)

---

## Struktur Folder

```
sharing-vision-api/
├── main.go
├── .env.example
├── go.mod
├── config/
│   └── db.go
├── enums/
│   └── status.go
├── models/
│   └── post.go
├── services/
│   └── post_service.go
├── controllers/
│   └── post_controller.go
├── routes/
│   └── routes.go
└── migration/
    └── 001_create_posts.sql
```

---

## Cara Menjalankan

### 1. Import Database

```bash
mysql -u root -p < migration/001_create_posts.sql
```

### 2. Setup Environment

```bash
cp .env.example .env
```

Isi file `.env` sesuai credentials MySQL kamu:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=sharing_vision
APP_PORT=8080
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Jalankan Server

#### Opsi A: Mode Standar

```bash
go run main.go
```

#### Opsi B: Mode Development (Auto-Reload menggunakan Air)

Sangat disarankan untuk development agar server otomatis restart saat ada perubahan kode.

1. Install Air (sekali saja):

```bash
go install github.com/air-verse/air@latest
```

2. Jalankan dengan Air:

```bash
air
```

Server berjalan di `http://localhost:8080`

---

## Endpoints

| Method   | URL                            | Deskripsi                        |
| -------- | ------------------------------ | -------------------------------- |
| `POST`   | `/article/`                    | Buat artikel baru                |
| `GET`    | `/article/list/:limit/:offset` | Ambil semua artikel (pagination) |
| `GET`    | `/article/:id`                 | Ambil artikel berdasarkan ID     |
| `PUT`    | `/article/:id`                 | Update artikel                   |
| `DELETE` | `/article/:id`                 | Hapus artikel                    |

---

## Request Payload (Create & Update)

```json
{
  "title": "string — minimal 20 karakter",
  "content": "string — minimal 200 karakter",
  "category": "string — minimal 3 karakter",
  "status": "publish | draft | thrash"
}
```

---

## Response Format

### Success

```json
{
  "status": "success",
  "message": "Post created successfully",
  "data": {}
}
```

> Field `data` hanya ada pada endpoint GET.

### Error

```json
{
  "status": "error",
  "message": "pesan error spesifik"
}
```

---

## HTTP Status Codes

| Code  | Kondisi                                |
| ----- | -------------------------------------- |
| `200` | Berhasil                               |
| `400` | Validasi gagal / parameter tidak valid |
| `404` | Data tidak ditemukan                   |
| `500` | Internal server error                  |

---

## Contoh Penggunaan (Postman)

**Create Post:**

```
POST http://localhost:8080/article/
Content-Type: application/json

{
  "title": "Judul artikel minimal 20 karakter",
  "content": "Isi konten yang harus panjang minimal dua ratus karakter. Tambahkan kalimat yang cukup agar validasi backend API lolos. Ini adalah contoh konten artikel yang dibuat untuk keperluan pengujian API menggunakan Postman.",
  "category": "Technology",
  "status": "publish"
}
```

**Get All Posts:**

```
GET http://localhost:8080/article/list/10/0
```

**Get by ID:**

```
GET http://localhost:8080/article/1
```

**Update Post:**

```
PUT http://localhost:8080/article/1
Content-Type: application/json

{
  "title": "Judul artikel minimal 20 karakter",
  "content": "Isi konten yang harus panjang minimal dua ratus karakter. Tambahkan kalimat yang cukup agar validasi backend API lolos. Ini adalah contoh konten artikel yang dibuat untuk keperluan pengujian API menggunakan Postman.",
  "category": "Technology",
  "status": "draft"
}
```

**Delete Post:**

```
DELETE http://localhost:8080/article/1
```

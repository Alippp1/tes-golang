# Tes Golang API

Backend API sederhana menggunakan **Golang + Fiber**.

Project ini dibuat untuk latihan / eksperimen REST API menggunakan Golang.

---

## 🚀 Tech Stack

- Golang
- Fiber
- GORM
- MySQL
- JWT (Authentication)

---

## ⚙️ Environment Variables

Untuk cara menggunakannya silahkan clone jika sudah lalu go mod tidy lalu untuk
menjalankannya go run cmd/server/main.go

Project ini menggunakan file `.env` untuk konfigurasi environment.  
File `.env` **tidak ikut di-push ke repository** (sudah di `.gitignore`).

### Contoh `.env`

Buat file `.env` di root project:

```env
APP_PORT=3000

DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=tes_purchasedb

JWT_SECRET=supersecretkey
WEBHOOK_URL=https://webhook.site/4a3608a4-e71e-43c4-b955-fdbfe893bc01

```

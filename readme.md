Berikut adalah contoh README.md yang lebih jelas dan mudah dipahami dengan langkah-langkah untuk menjalankan proyek Anda:

```markdown
# Go Rental

Proyek ini adalah aplikasi manajemen rental yang dibangun menggunakan bahasa pemrograman Go dan PostgreSQL sebagai basis datanya.

## Struktur Proyek
```
```
go-rental/
├── db/
│   ├── migration-golang/
│   │   ├── 20231205131104_create_users_table.up.sql
│   │   ├── 20231205131104_create_users_table.down.sql
│   │   └── ...
├── main.go
├── models/
│   ├── user.go
│   └── ...
├── controllers/
│   ├── userController.go
│   └── ...
├── routes/
│   ├── userRoutes.go
│   └── ...
├── config/
│   ├── config.go
│   └── ...
└── README.md
```

## Deskripsi

Aplikasi ini menyediakan fitur-fitur untuk mengelola data rental, seperti menambah, mengubah, menghapus, dan melihat data rental. Aplikasi ini menggunakan migrasi database untuk mengelola skema database.

## Langkah-langkah Menjalankan Proyek

### Prasyarat
1. **Golang**: Pastikan Golang sudah terpasang di sistem Anda. Anda bisa mengunduhnya dari [situs resmi Golang](https://golang.org/dl/).
2. **PostgreSQL**: Pastikan PostgreSQL sudah terpasang dan dikonfigurasi dengan benar di sistem Anda. Anda bisa mengunduhnya dari [situs resmi PostgreSQL](https://www.postgresql.org/download/).
3. **Migrasi Database**: Menggunakan [golang-migrate](https://github.com/golang-migrate/migrate) untuk mengelola migrasi database.

### Mengkonfigurasi Proyek

1. **Mengatur Database**:
   - Buat database baru di PostgreSQL untuk proyek ini.
   - Konfigurasi detail koneksi database di file `config/config.go`.

2. **Menjalankan Migrasi Database**:
   - Pastikan Anda berada di direktori proyek `go-rental/`.
   - Jalankan perintah berikut untuk menjalankan migrasi database:
     ```sh
     migrate -path ./db/migration-golang -database postgres://postgres:mysecretpassword@localhost:5432/Rental?sslmode=disable up
     ```
   - Jika perlu memaksa migrasi ke versi tertentu, gunakan perintah berikut:
     ```sh
     migrate -path ./db/migration-golang -database postgres://postgres:mysecretpassword@localhost:5432/Rental?sslmode=disable force 20231205131104
     ```

### Menjalankan Aplikasi

1. **Menjalankan Aplikasi**:
   - Jalankan perintah berikut untuk memulai aplikasi:
     ```sh
     go run cmd/server/main.go
     ```
   - Aplikasi akan berjalan di port yang dikonfigurasi (default: `8080`).

### Struktur Direktori

- **db/migration-golang/**: Berisi file-file migrasi untuk mengelola skema database.
- **main.go**: File utama untuk menjalankan aplikasi.
- **models/**: Berisi definisi model untuk entitas dalam aplikasi.
- **controllers/**: Berisi logika bisnis dan pengendali untuk setiap entitas.
- **routes/**: Berisi definisi rute untuk API.
- **config/**: Berisi konfigurasi aplikasi.

### Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan buat pull request atau buka issue untuk diskusi lebih lanjut.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
```

Semoga ini membantu! Jika ada pertanyaan atau perlu penjelasan lebih lanjut, jangan ragu untuk menghubungi.
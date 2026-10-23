# To-Do List REST API (Golang)

Aplikasi backend RESTful API sederhana untuk pencatatan dan pengelolaan daftar tugas (*To-Do List*), dibangun menggunakan bahasa pemrograman **Go** dengan memanfaatkan **Go Standard Library** (terutama `net/http`) tanpa framework atau dependensi pihak ketiga, serta menggunakan penyimpanan data sementara di memori (*in-memory slice*).

---

## Daftar Isi

- [Gambaran Proyek](#gambaran-proyek)
- [Fitur Utama](#fitur-utama)
- [Struktur Direktori](#struktur-direktori)
- [Prasyarat Sistem](#prasyarat-sistem)
- [Cara Menjalankan](#cara-menjalankan)
  - [Menjalankan Server Langsung](#menjalankan-server-langsung)
  - [Melakukan Build Binary](#melakukan-build-binary)
- [Menjalankan Pengujian (Unit Testing)](#menjalankan-pengujian-unit-testing)
- [Dokumentasi API (Endpoints)](#dokumentasi-api-endpoints)
  - [1. Tambah Tugas Baru](#1-tambah-tugas-baru)
  - [2. Lihat Seluruh Daftar Tugas](#2-lihat-seluruh-daftar-tugas)
  - [3. Tandai Tugas Selesai](#3-tandai-tugas-selesai)
- [Contoh Pengujian dengan cURL](#contoh-pengujian-dengan-curl)
  - [Linux / macOS / Git Bash](#linux--macos--git-bash)
  - [Windows PowerShell](#windows-powershell)
- [Panduan Pengujian dengan Postman](#panduan-pengujian-dengan-postman)
- [Kode Status HTTP & Penanganan Error](#kode-status-http--penanganan-error)

---

## Gambaran Proyek

Proyek ini dirancang dengan arsitektur modular yang memisahkan tanggung jawab antara lapisan data (*database/storage*), lapisan logika HTTP (*handlers & models*), dan konfigurasi perutean (*router*):

- **Zero External Dependencies**: Menggunakan library bawaan Go (`net/http`, `encoding/json`, `sync`, dll.).
- **Go 1.22+ Enhanced Routing**: Memanfaatkan fitur bawaan `http.ServeMux` untuk penanganan HTTP method dan path parameter (`GET /tasks`, `/tasks/{id}/done`).
- **In-Memory Storage**: Data tugas disimpan di dalam slice dengan perlindungan *thread-safe* (`sync.Mutex`).

---

## Fitur Utama

1. **Tambah Tugas (`Create Task`)**: Menambahkan tugas baru ke dalam daftar. Otomatis memberikan `id` unik dan status awal `done = false`.
2. **Lihat Daftar Tugas (`Get Tasks`)**: Menampilkan seluruh data tugas yang tersimpan dalam format JSON array.
3. **Tandai Selesai (`Mark Task Done`)**: Memperbarui status tugas tertentu menjadi selesai (`done = true`) berdasarkan ID tugas.
4. **Validasi & Error Handling**:
   - Menolak input jika judul tugas kosong atau hanya berisi spasi (`400 Bad Request`).
   - Memberikan respon jika ID tugas tidak valid atau tidak ditemukan (`404 Not Found`).

---

## Struktur Direktori

```text
golang_blkpp/
├── database/
│   └── database.go    # In-memory storage slice & thread-safe operation (sync.Mutex)
├── router/
│   └── router.go      # Konfigurasi ServeMux dan pendaftaran rute API
├── todo/
│   ├── handler.go     # HTTP Handler & logika request/response JSON
│   ├── model.go       # Definisi struct Task dan interface Store
│   └── routes.go      # Pemetaan URL endpoint ke fungsi handler
├── go.mod             # File definisi modul Go
├── main.go            # Titik masuk utama (entry point) aplikasi
├── main_test.go       # Unit test & integration test untuk endpoint API
├── prompt.md          # Spesifikasi kebutuhan & kriteria aplikasi
└── README.md          # Dokumentasi proyek
```

---

## Prasyarat Sistem

- **Go (Golang)**: Versi `1.22.0` ke atas (disarankan sesuai versi pada `go.mod`).
- Terminal / Command Prompt / PowerShell / Bash.
- Tools pengujian HTTP (cURL, Postman, Thunder Client, dll.).

Periksa versi Go yang terpasang dengan perintah:
```bash
go version
```

---

## Cara Menjalankan

### Menjalankan Server Langsung

Jalankan perintah berikut di direktori root proyek:

```bash
go run main.go
```

Jika berhasil, server akan berjalan di:
```text
Server To-Do List berjalan di http://localhost:8080
```

### Melakukan Build Binary

Untuk mengompilasi program menjadi file binary mandiri (*executable*):

**Linux / macOS:**
```bash
go build -o todo-app main.go
./todo-app
```

**Windows:**
```powershell
go build -o todo-app.exe main.go
.\todo-app.exe
```

---

## Menjalankan Pengujian (Unit Testing)

Aplikasi dilengkapi unit test lengkap di `main_test.go` yang menguji skenario sukses maupun skenario error (*edge cases*).

Jalankan seluruh test dengan perintah:

```bash
go test -v ./...
```

---

## Dokumentasi API (Endpoints)

Base URL default: `http://localhost:8080`

### 1. Tambah Tugas Baru

Menambahkan catatan tugas baru.

- **URL:** `/tasks`
- **Method:** `POST`
- **Headers:** `Content-Type: application/json`
- **Request Body:**
  ```json
  {
    "title": "Belajar Golang"
  }
  ```
- **Responses:**
  - **`201 Created`** (Berhasil dibuat):
    ```json
    {
      "id": 1,
      "title": "Belajar Golang",
      "done": false
    }
    ```
  - **`400 Bad Request`** (Judul kosong atau format JSON salah):
    ```json
    {
      "error": "judul tugas tidak boleh kosong"
    }
    ```

---

### 2. Lihat Seluruh Daftar Tugas

Mendapatkan seluruh daftar tugas yang sedang tersimpan di memori.

- **URL:** `/tasks`
- **Method:** `GET`
- **Headers:** None
- **Responses:**
  - **`200 OK`**:
    ```json
    [
      {
        "id": 1,
        "title": "Belajar Golang",
        "done": false
      }
    ]
    ```
    *Catatan: Mengembalikan array kosong `[]` jika belum ada tugas.*

---

### 3. Tandai Tugas Selesai

Mengubah status `done` menjadi `true` berdasarkan ID tugas. Endpoint ini mendukung metode `PATCH` maupun `PUT`.

- **URL:** `/tasks/{id}/done`
- **Method:** `PATCH` atau `PUT`
- **URL Parameter:** `id` (integer) — ID tugas yang ingin diubah.
- **Responses:**
  - **`200 OK`** (Berhasil diubah):
    ```json
    {
      "id": 1,
      "title": "Belajar Golang",
      "done": true
    }
    ```
  - **`400 Bad Request`** (Format ID bukan angka):
    ```json
    {
      "error": "ID tugas tidak valid"
    }
    ```
  - **`404 Not Found`** (ID tugas tidak ditemukan di memori):
    ```json
    {
      "error": "tugas tidak ditemukan"
    }
    ```

---

## Contoh Pengujian dengan cURL

Berikut contoh perintah cURL yang dapat langsung dijalankan di terminal:

### Linux / macOS / Git Bash

```bash
# 1. Tambah Tugas 1
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar REST API Golang"}'

# 2. Tambah Tugas 2
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Membaca dokumentasi Go"}'

# 3. Lihat Semua Tugas
curl -X GET http://localhost:8080/tasks

# 4. Tandai Tugas ID 1 Selesai (PATCH atau PUT)
curl -X PATCH http://localhost:8080/tasks/1/done

# 5. Uji Validasi Judul Kosong (Harus mengembalikan 400 Bad Request)
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "   "}'

# 6. Uji ID Tidak Ditemukan (Harus mengembalikan 404 Not Found)
curl -X PATCH http://localhost:8080/tasks/999/done
```

### Windows PowerShell

Jika menggunakan PowerShell, gunakan sintaks berikut:

```powershell
# 1. Tambah Tugas Baru
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/tasks" -ContentType "application/json" -Body '{"title": "Belajar REST API Golang"}'

# 2. Lihat Semua Tugas
Invoke-RestMethod -Method Get -Uri "http://localhost:8080/tasks"

# 3. Tandai Tugas Selesai
Invoke-RestMethod -Method Patch -Uri "http://localhost:8080/tasks/1/done"
```

*Atau menggunakan `curl.exe` bawaan Windows:*
```cmd
curl.exe -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\": \"Belajar REST API Golang\"}"
curl.exe -X GET http://localhost:8080/tasks
curl.exe -X PATCH http://localhost:8080/tasks/1/done
```

---

## Panduan Pengujian dengan Postman

1. **Membuat Request Baru**:
   - Buka Postman dan buat Collection baru, misalnya `To-Do App`.
2. **POST /tasks (Tambah Tugas)**:
   - Pilih Method: `POST`
   - URL: `http://localhost:8080/tasks`
   - Tab **Body** > pilih **raw** > pilih tipe **JSON**
   - Masukkan JSON:
     ```json
     {
       "title": "Mengerjakan tugas BLK"
     }
     ```
   - Klik **Send** dan pastikan status `201 Created`.
3. **GET /tasks (Lihat Daftar)**:
   - Pilih Method: `GET`
   - URL: `http://localhost:8080/tasks`
   - Klik **Send** dan pastikan status `200 OK` dengan daftar tugas.
4. **PATCH /tasks/{id}/done (Tandai Selesai)**:
   - Pilih Method: `PATCH`
   - URL: `http://localhost:8080/tasks/1/done`
   - Klik **Send** dan pastikan status `200 OK` dengan `"done": true`.

---

## Kode Status HTTP & Penanganan Error

| Status Code | Kondisi | Keterangan |
| :--- | :--- | :--- |
| `200 OK` | Sukses `GET` / `PATCH` / `PUT` | Data berhasil diambil atau status berhasil diperbarui. |
| `201 Created` | Sukses `POST` | Tugas baru berhasil disimpan ke dalam memori. |
| `400 Bad Request` | Validasi Gagal | Format JSON salah, title kosong/hanya spasi, atau ID bukan angka. |
| `404 Not Found` | Data Tidak Ditemukan | ID tugas yang diminta tidak terdapat dalam memori. |

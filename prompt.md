# Prompt Pembuatan Aplikasi To-Do List

Buatkan aplikasi **To-Do List** menggunakan bahasa pemrograman **Go**
sesuai dengan tujuan, fitur, teknologi, data, dan kriteria keberhasilan
berikut.

## 1. Tujuan Aplikasi

Saya ingin membuat aplikasi **To-Do List** untuk mencatat dan mengelola
daftar tugas.

## 2. Pengguna dan Masalah

-   **Pengguna:** Saya sendiri.
-   **Masalah yang ingin diselesaikan:** Membantu mengingat dan
    menyelesaikan tugas agar tidak terlewat.

## 3. Tiga Fitur Wajib (Must-Have)

Aplikasi wajib memiliki tiga fitur berikut:

1.  **Tambah tugas** --- pengguna dapat menambahkan tugas baru.
2.  **Lihat daftar** --- pengguna dapat melihat semua tugas yang
    tersimpan.
3.  **Tandai selesai** --- pengguna dapat mengubah status tugas menjadi
    selesai.

## 4. Teknologi dan Batasan

-   **Bahasa backend:** Go, menggunakan standard library, terutama
    `net/http`.
-   **Penyimpanan data:** Simpan data di memori menggunakan `slice`
    (in-memory slice).
-   Jangan menggunakan database atau framework web eksternal untuk versi
    ini.

## 5. Data dan API

Gunakan struktur data `Task` dengan field berikut:

-   `ID` --- identitas unik tugas.
-   `Title` --- judul tugas.
-   `Done` --- status selesai (`true` atau `false`).

Buat REST API yang sesuai dengan kebutuhan aplikasi. Minimal, sediakan
endpoint untuk:

-   Menambahkan tugas.
-   Melihat seluruh daftar tugas.
-   Menandai tugas sebagai selesai berdasarkan ID.

Gunakan format JSON untuk request dan response API. Tentukan metode HTTP
dan URL endpoint dengan jelas, lalu jelaskan cara menggunakannya.

## 6. Kriteria Berhasil (Acceptance Criteria)

Aplikasi dinyatakan berhasil jika memenuhi semua kriteria berikut:

-   [ ] Program berhasil di-build (`go build`).
-   [ ] Mengembalikan HTTP **400 Bad Request** jika judul tugas kosong.
-   [ ] Mengembalikan HTTP **404 Not Found** jika ID tugas tidak
    ditemukan.
-   [ ] Fitur tambah tugas, lihat daftar tugas, dan tandai tugas selesai
    dapat digunakan.
-   [ ] Data disimpan dalam `slice` selama aplikasi berjalan.
-   [ ] API menerima dan mengembalikan data dalam format JSON.

## Instruksi Hasil

Berikan hasil berupa:

1.  Kode program Go lengkap yang dapat langsung dijalankan.
2.  Struktur folder dan nama file yang digunakan.
3.  Penjelasan singkat setiap endpoint, metode HTTP, request, dan
    response.
4.  Contoh pengujian menggunakan `curl` atau Postman.
5.  Perintah untuk menjalankan dan melakukan build program.

Pastikan kode mudah dipahami oleh pemula, memiliki penanganan error yang
sesuai, dan memenuhi seluruh kriteria keberhasilan di atas.

package todo

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler membungkus Store untuk menangani request HTTP
type Handler struct {
	store Store
}

// NewHandler membuat instance Handler baru dengan Store yang diberikan
func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

// Helper untuk mengirim response JSON
func sendJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Gagal melakukan encode JSON: %v", err)
	}
}

// Helper untuk mengirim response error JSON
func sendError(w http.ResponseWriter, status int, message string) {
	sendJSON(w, status, map[string]string{"error": message})
}

// GetTasks menangani request GET /tasks (Melihat seluruh daftar tugas)
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// 1. Ambil seluruh data tugas dari store (h.store.GetAll())
	// 2. Kirim respons JSON dengan status HTTP 200 OK
}

// CreateTask menangani request POST /tasks (Menambahkan tugas baru)
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// 1. Baca request body JSON (misal struct dengan field Title)
	// 2. Validasi judul tugas: jika kosong/hanya spasi, kembalikan HTTP 400 Bad Request
	// 3. Simpan tugas baru ke store (h.store.Create(title))
	// 4. Kirim respons JSON tugas yang berhasil dibuat dengan status HTTP 201 Created
}

// MarkTaskDone menangani request PATCH/PUT /tasks/{id}/done (Menandai selesai)
func (h *Handler) MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// 1. Ambil ID dari path URL menggunakan r.PathValue("id") dan konversi ke integer
	// 2. Jika ID tidak valid/bukan angka, kembalikan HTTP 400 Bad Request
	// 3. Panggil store untuk menandai selesai (h.store.MarkDone(id))
	// 4. Jika ID tidak ditemukan, kembalikan HTTP 404 Not Found
	// 5. Jika berhasil, kirim respons JSON tugas yang diperbarui dengan status HTTP 200 OK
}

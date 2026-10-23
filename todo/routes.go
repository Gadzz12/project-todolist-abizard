package todo

import "net/http"

// RegisterRoutes mendaftarkan seluruh endpoint untuk fitur Todo ke ServeMux
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /tasks", h.GetTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	mux.HandleFunc("PATCH /tasks/{id}/done", h.MarkTaskDone)
	mux.HandleFunc("PUT /tasks/{id}/done", h.MarkTaskDone)
}

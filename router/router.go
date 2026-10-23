package router

import (
	"net/http"
	"todo-app/todo"
)

// Setup mengonfigurasi dan mengembalikan ServeMux untuk seluruh rute API
func Setup(h *todo.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	return mux
}

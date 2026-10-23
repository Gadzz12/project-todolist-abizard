package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/database"
	"todo-app/router"
	"todo-app/todo"
)

func main() {
	// Inisialisasi Database (in-memory)
	db := database.NewDB()

	// Inisialisasi Handler Todo
	todoHandler := todo.NewHandler(db)

	// Inisialisasi Router
	r := router.Setup(todoHandler)

	port := ":8080"
	fmt.Printf("Server To-Do List berjalan di http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server gagal berjalan: %v", err)
	}
}

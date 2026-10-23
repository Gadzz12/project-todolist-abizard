package database

import (
	"sync"
	"todo-app/todo"
)

// DB mengelola penyimpanan data tugas di memori (in-memory slice)
type DB struct {
	tasks  []todo.Task
	nextID int
	mu     sync.Mutex
}

// NewDB menginisialisasi instance DB baru
func NewDB() *DB {
	return &DB{
		tasks:  []todo.Task{},
		nextID: 1,
	}
}

// GetAll mengambil semua tugas yang tersimpan
func (db *DB) GetAll() []todo.Task {
	// TODO: Implementasikan logika untuk mengambil seluruh data tugas dari slice
	return nil
}

// Create menambahkan tugas baru ke dalam slice
func (db *DB) Create(title string) todo.Task {
	// TODO: Implementasikan logika untuk menambahkan tugas baru ke dalam slice
	// Pastikan ID auto-increment dan thread-safe
	return todo.Task{}
}

// MarkDone memperbarui status tugas menjadi selesai berdasarkan ID
func (db *DB) MarkDone(id int) (todo.Task, bool) {
	// TODO: Implementasikan logika untuk mencari tugas berdasarkan ID
	// dan mengubah nilai Done menjadi true
	return todo.Task{}, false
}

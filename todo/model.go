package todo

// Task merepresentasikan struktur data tugas sesuai spesifikasi
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Store adalah interface untuk abstraksi operasi data tugas
type Store interface {
	GetAll() []Task
	Create(title string) Task
	MarkDone(id int) (Task, bool)
}

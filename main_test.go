package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/database"
	"todo-app/router"
	"todo-app/todo"
)

func setupTestServer() http.Handler {
	db := database.NewDB()
	h := todo.NewHandler(db)
	return router.Setup(h)
}

func TestTodoListAPI(t *testing.T) {
	handler := setupTestServer()

	// 1. Test GET /tasks saat kosong
	t.Run("GET /tasks empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %d", rec.Code)
		}

		var list []todo.Task
		if err := json.Unmarshal(rec.Body.Bytes(), &list); err != nil {
			t.Fatalf("Failed to decode JSON: %v", err)
		}
		if len(list) != 0 {
			t.Fatalf("Expected 0 tasks, got %d", len(list))
		}
	})

	// 2. Test POST /tasks dengan title kosong -> 400 Bad Request
	t.Run("POST /tasks empty title (400 Bad Request)", func(t *testing.T) {
		payload := []byte(`{"title": "   "}`)
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("Expected status 400, got %d", rec.Code)
		}
	})

	// 3. Test POST /tasks sukses -> 201 Created
	t.Run("POST /tasks success", func(t *testing.T) {
		payload := []byte(`{"title": "Belajar Golang"}`)
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Fatalf("Expected status 201, got %d", rec.Code)
		}

		var created todo.Task
		if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
			t.Fatalf("Failed to decode JSON: %v", err)
		}

		if created.ID != 1 || created.Title != "Belajar Golang" || created.Done != false {
			t.Fatalf("Unexpected created task: %+v", created)
		}
	})

	// 4. Test GET /tasks setelah penambahan
	t.Run("GET /tasks contains created task", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %d", rec.Code)
		}

		var list []todo.Task
		if err := json.Unmarshal(rec.Body.Bytes(), &list); err != nil {
			t.Fatalf("Failed to decode JSON: %v", err)
		}
		if len(list) != 1 {
			t.Fatalf("Expected 1 task, got %d", len(list))
		}
		if list[0].Title != "Belajar Golang" {
			t.Fatalf("Expected title 'Belajar Golang', got '%s'", list[0].Title)
		}
	})

	// 5. Test PATCH /tasks/{id}/done untuk ID yang tidak ada -> 404 Not Found
	t.Run("PATCH /tasks/999/done (404 Not Found)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/tasks/999/done", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("Expected status 404, got %d", rec.Code)
		}
	})

	// 6. Test PATCH /tasks/1/done sukses -> 200 OK & Done = true
	t.Run("PATCH /tasks/1/done success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/tasks/1/done", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %d", rec.Code)
		}

		var updated todo.Task
		if err := json.Unmarshal(rec.Body.Bytes(), &updated); err != nil {
			t.Fatalf("Failed to decode JSON: %v", err)
		}

		if !updated.Done {
			t.Fatalf("Expected task to be done, got false")
		}
	})
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	tasks  = []Task{}
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)
	log.Println("Api en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mu.Lock()
		json.NewEncoder(w).Encode(tasks)
		mu.Unlock()
	case "POST":
		var t Task
		json.NewDecoder(r.Body).Decode(&t)
		mu.Lock()
		t.ID = nextID
		nextID++
		t.CreatedAt = time.Now()
		tasks = append(tasks, t)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.ID == id {
			switch r.Method {
			case "GET":
				json.NewEncoder(w).Encode(t)

			case "PUT":
				json.NewDecoder(r.Body).Decode(&t)
				t.ID = id
				tasks[i] = t
				json.NewEncoder(w).Encode(t)
			case "DELETE":
				tasks = append(tasks[:i], tasks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
			}
			return
		}
	}
	http.NotFound(w, r)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"snippets/workshops/workshop-01/handlers"
	"snippets/workshops/workshop-01/types"
)

var (
	info types.Displayer = types.InitializeInfo()
	data                 = &types.Data{
		Tasks:  make(map[int]types.Task),
		NextID: 1,
	}
)

func init() {
	// Add some initial tasks
	data.Tasks[data.NextID] = types.Task{ID: data.NextID, Title: "Learn Go basics", Completed: false}
	data.NextID++
	data.Tasks[data.NextID] = types.Task{ID: data.NextID, Title: "Build a simple API", Completed: false}
	data.NextID++
}

func main() {
	fmt.Println("Starting simple task API...")
	info.Display()

	// Display tasks using the interface
	fmt.Println("Initial tasks:")
	for _, task := range data.Tasks {
		// Using task as TaskDisplayer interface
		var displayer types.Displayer = task
		displayer.Display()
	}

	// Start a goroutine for periodic task printing
	done := make(chan bool)
	go logTasksPeriodically(data, done)

	// Set up API routes
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.HandleGetTasks(w, data)
		case http.MethodPost:
			handlers.HandleCreateTask(w, r, data)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.HandleGetTask(w, r, data)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		if err != http.ErrServerClosed {
			done <- true
			fmt.Printf("Server crashed with error: %v\n", err)
			os.Exit(1)
		}
	}
}

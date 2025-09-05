package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snippets/workshops/workshop-01/types"
	"strconv"
)

// HandleGetTasks returns all tasks
func HandleGetTasks(w http.ResponseWriter, data *types.Data) {
	// Create a slice to hold all tasks
	allTasks := []types.Task{}

	// Iterate through the map to collect all tasks
	for _, task := range data.Tasks {
		allTasks = append(allTasks, task)
	}

	// Set content type and encode response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allTasks)
}

// HandleCreateTask creates a new task
func HandleCreateTask(w http.ResponseWriter, r *http.Request, data *types.Data) {
	// Create a new task from request body
	var task types.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	// Create a channel for the result
	resultChan := make(chan types.Task)

	// Process task in a goroutine
	go func() {
		// Assign ID and save task
		task.ID = data.NextID
		data.Tasks[data.NextID] = task
		data.NextID++

		// Send task to channel
		resultChan <- task
	}()

	// Wait for task to be processed
	createdTask := <-resultChan

	// Respond with created task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

// HandleGetTask gets a single task by ID
func HandleGetTask(w http.ResponseWriter, r *http.Request, data *types.Data) {
	// Using defer for demonstration
	defer fmt.Println("Task retrieval completed")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	// Check for invalid ID
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Find task by ID
	task, exists := data.Tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Return the task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

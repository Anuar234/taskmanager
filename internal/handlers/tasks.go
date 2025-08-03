package handlers

import (
	"encoding/json"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []string{"Task-1", "Task-2", "Task-3"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
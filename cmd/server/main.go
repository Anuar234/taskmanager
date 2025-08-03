package main

import (
	"log"
	"net/http"

	"taskmanager/internal/database"
	"taskmanager/internal/handlers"
)

func main() {

	database.ConnectMongo()

	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetTasks(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateTask(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetTasksHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateTaskHandler(w, r)
		} else {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/tasks/create", handlers.CreateTaskHandler)
	http.HandleFunc("/api/tasks/update/", handlers.UpdateTaskHandler)
	http.HandleFunc("/api/tasks/delete/", handlers.DeleteTaskHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

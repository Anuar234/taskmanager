package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"taskmanager/internal/models"
	"taskmanager/internal/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []string{"Task-1", "Task-2", "Task-3"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Ошибка при чтении тела запроса", http.StatusBadRequest)
		return
	}

	task.ID = primitive.NewObjectID()
	coll := mongodb.GetTaskCollection()
	_, err := coll.InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, "Ошибка при создании задачи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

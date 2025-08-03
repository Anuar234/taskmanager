package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/internal/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection = database.GetCollection("tasks")

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}

	task.Done = false

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := taskCollection.InsertOne(ctx, task)
	if err != nil {
		http.Error(w, "failed to insert task", http.StatusInternalServerError)
		return
	}

	task.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "failed tp get tasks", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			http.Error(w, "error decosing task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/update")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Error duirng decoding", http.StatusBadRequest)
		return
	}

	coll := mongodb.GetTaskCollection()
	update := bson.M{
		"$set": bson.M{
			"title":     updatedTask.Title,
			"completed": updatedTask.Done,
		},
	}

	_, err = coll.UpdateByID(context.TODO(), id, update)
	if err != nil {
		http.Error(w, "Error duirng update", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task updatded")
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/tasks/delete")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	coll := mongodb.GetTaskCollection()
	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, "error during deleting", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task deleted")
}

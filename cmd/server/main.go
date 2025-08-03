package main

import (
	"fmt"
	"log"
	"net/http"

	"taskmanager/internal/handlers"
	"taskmanager/internal/mongodb"
)

func main() {
    // Подключение к Mongo
    coll := mongodb.GetTaskCollection()
    fmt.Println("✅ MongoDB подключён. Коллекция:", coll.Name())

    http.HandleFunc("/api/tasks", handlers.GetTasksHandler)


    fmt.Println("🚀 Сервер запущен на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


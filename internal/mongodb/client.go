package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance    *mongo.Client
    clientInstanceErr error
    mongoOnce         sync.Once
)

const (
    dbName         = "taskmanager"
    collectionName = "tasks"
)

func Connect() (*mongo.Client, error) {
    mongoOnce.Do(func() {
        uri := os.Getenv("MONGO_URI")
        if uri == "" {
            uri = "mongodb://localhost:27017"
        }

        clientOptions := options.Client().ApplyURI(uri)

        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        client, err := mongo.Connect(ctx, clientOptions)
        if err != nil {
            clientInstanceErr = err
            return
        }

        err = client.Ping(ctx, nil)
        if err != nil {
            clientInstanceErr = err
            return
        }

        clientInstance = client
        fmt.Println("✅ MongoDB connected")
    })

    return clientInstance, clientInstanceErr
}

func GetTaskCollection() *mongo.Collection {
    client, err := Connect()
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    return client.Database(dbName).Collection(collectionName)
}

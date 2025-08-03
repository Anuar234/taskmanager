package mongodb

import (
	"context"
	"taskmanager/internal/models"
	"time"
)

func InsertTask(task models.Task) error {
	coll := GetTaskCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := coll.InsertOne(ctx, task)
	return err
}
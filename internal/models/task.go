package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omniempty" json:"id"`
	Title     string             `bson:"title" json:"desciption"`
	Done      bool               `bson:"done" json:"done"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

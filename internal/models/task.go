package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID primitive.ObjectID `bson:"_id,omniempty" json:"id,omniempty"`
	Title string `bson:"title" json:"title"`
}
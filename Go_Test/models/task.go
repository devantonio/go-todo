package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task represents a task entity in the application.
type Task struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Title     string             `bson:"title"`
    Completed bool               `bson:"completed"`
    Description string           `bson:"description"`
}

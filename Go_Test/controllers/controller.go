package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskController handles requests related to tasks.
type TaskController struct {
	DB *mongo.Collection
}

// NewTaskController creates a new TaskController with the given database client.
func NewTaskController(client *mongo.Client, dbName string) *TaskController {
	collection := client.Database(dbName).Collection("tasks")
	return &TaskController{DB: collection}
}

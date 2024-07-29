package controllers

import (
    "context"
    "encoding/json"
    "net/http"
	"log"
	"go_test/utils"
    "go_test/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive" // Add this import
)

// TaskController handles requests related to tasks.
type TaskController struct {
    DB *mongo.Collection
}

// GetTasks handles GET requests to fetch all tasks.
func (tc *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []models.Task
    cursor, err := tc.DB.Find(context.Background(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var task models.Task
        cursor.Decode(&task)
        tasks = append(tasks, task)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

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

// NewTaskController creates a new TaskController with the given database client.
func NewTaskController(client *mongo.Client, dbName string) *TaskController {
    collection := client.Database(dbName).Collection("tasks")
    return &TaskController{DB: collection}
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
    // Use the ReadAndLogBody function to read and log the request body
    body, err := utils.ReadBody(r)

	// Log the body content
	log.Println("Body:", string(body))

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Decode the task from the body
    var task models.Task
    if err := json.Unmarshal(body, &task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    task.ID = primitive.NewObjectID() // Use primitive.NewObjectID
    _, err = tc.DB.InsertOne(context.Background(), task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

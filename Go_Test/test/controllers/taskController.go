// controllers/taskController.go
package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "../models"
)

// Global variable to hold the MongoDB client
var client *mongo.Client

// SetMongoClient sets the MongoDB client for use in the controller
func SetMongoClient(c *mongo.Client) {
    client = c
}

// getTasks handles GET requests to retrieve all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var tasks []models.Task

    // Get a handle to the tasks collection
    collection := client.Database("taskdb").Collection("tasks")

    // Set a timeout for the operation
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find all tasks in the collection
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }
    defer cursor.Close(ctx)

    // Iterate through the cursor and decode each document
    for cursor.Next(ctx) {
        var task models.Task
        cursor.Decode(&task)
        tasks = append(tasks, task)
    }

    // Check if the cursor encountered any errors
    if err := cursor.Err(); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    // Encode and return the tasks
    json.NewEncoder(w).Encode(tasks)
}

// createTask handles POST requests to create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var task models.Task

    // Decode the incoming task json
    json.NewDecoder(r.Body).Decode(&task)

    // Get a handle to the tasks collection
    collection := client.Database("taskdb").Collection("tasks")

    // Set a timeout for the operation
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Insert the new task into the collection
    result, err := collection.InsertOne(ctx, task)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    // Return the result
    json.NewEncoder(w).Encode(result)
}

// getTask handles GET requests to retrieve a specific task by ID
func getTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Get the task ID from the URL parameters
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var task models.Task

    // Get a handle to the tasks collection
    collection := client.Database("taskdb").Collection("tasks")

    // Set a timeout for the operation
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find the task with the matching ID
    err := collection.FindOne(ctx, models.Task{ID: id}).Decode(&task)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    // Encode and return the task
    json.NewEncoder(w).Encode(task)
}

// updateTask handles PUT requests to update a specific task
func updateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Get the task ID from the URL parameters
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var task models.Task

    // Decode the incoming task json
    json.NewDecoder(r.Body).Decode(&task)

    // Get a handle to the tasks collection
    collection := client.Database("taskdb").Collection("tasks")

    // Set a timeout for the operation
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Create an update document
    update := bson.M{
        "$set": bson.M{
            "title":       task.Title,
            "description": task.Description,
            "completed":   task.Completed,
        },
    }

    // Update the task in the collection
    _, err := collection.UpdateOne(ctx, models.Task{ID: id}, update)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    // Encode and return the updated task
    json.NewEncoder(w).Encode(task)
}

// deleteTask handles DELETE requests to remove a specific task
func deleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Get the task ID from the URL parameters
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    // Get a handle to the tasks collection
    collection := client.Database("taskdb").Collection("tasks")

    // Set a timeout for the operation
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Delete the task from the collection
    _, err := collection.DeleteOne(ctx, models.Task{ID: id})
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "` + err.Error() + `"}`))
        return
    }

    // Return a success message
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Task deleted successfully"}`))
}
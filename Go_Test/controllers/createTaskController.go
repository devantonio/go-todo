package controllers

import (
	"context"
	"encoding/json"
	"go_test/models"
	"go_test/utils"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive" // Add this import
)

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Use the ReadAndLogBody function to read the request body from the client
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

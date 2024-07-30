package controllers

import (
	"context"
	"encoding/json"
	"go_test/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetTasks handles GET requests to fetch all tasks.
func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
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

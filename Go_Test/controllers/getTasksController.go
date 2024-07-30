package controllers

import (
	"context"
	"go_test/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetTasks handles GET requests to fetch all tasks as an API response.
func (tc *TaskController) GetTasks(limit int64) ([]models.Task, error) {
	var tasks []models.Task

	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(limit)
	}
	cursor, err := tc.DB.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

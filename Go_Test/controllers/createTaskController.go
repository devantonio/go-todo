package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go_test/models"
	"go_test/utils"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive" // Add this import
)

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Use the ReadAndLogBody function to read the request body from the client
	body, err := utils.ReadBody(r)

	// Log the body content from the client
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

	// Request Spotify API
	response, err := spotifyApiRequest()
	if err != nil {
		fmt.Println("Request failed while creating task:", err)
		return
	}

	fmt.Println(response)

	// Marshal the response into JSON
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// send the task to an email
	//utils.SendEmail("Go To", "btrflyefct@outlook.com", "A Task Was Created", "Congrats on creating your first task")
	featuredCommentsResponse, err := getFeaturedComments()
	if err != nil {
		fmt.Println("Request failed while getting featured comments:", err)
		return
	}

	// Marshal the response into JSON
	jsonComments, err := json.Marshal(featuredCommentsResponse)
	if err != nil {
		fmt.Println("whoops:", err)
		return
	}
	// handle featured comments

	// write status and created task to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonComments)
	//json.NewEncoder(w).Encode(task)
	//json.NewEncoder(w).Encode(jsonResponse)
}

package routes

import (
	"go_test/controllers"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the routes for the application.
func RegisterRoutes(router *mux.Router, taskController *controllers.TaskController) {
	// Serve static files from the /static directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fetch data from the database or any other source
		tasks, err := taskController.GetTasks(2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse and execute the template
		tmpl, err := template.ParseFiles("static/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, tasks)
	})

	// API routes
	router.HandleFunc("/tasks", taskController.GetTask).Methods("GET")
	router.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
}

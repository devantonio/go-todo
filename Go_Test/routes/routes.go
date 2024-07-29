package routes

import (
    "net/http"
	"log"
    "github.com/gorilla/mux"
    "go_test/controllers"
)

// RegisterRoutes sets up the routes for the application.
func RegisterRoutes(router *mux.Router, taskController *controllers.TaskController) {
    // Serve static files from the /static directory
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving index.html")
		http.ServeFile(w, r, "static/templates/index.html")
	})
	
    // API routes
    router.HandleFunc("/tasks", taskController.GetTasks).Methods("GET")
    router.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
}

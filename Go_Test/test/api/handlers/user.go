package handlers

import (
    "net/http"
    "appname/internal/services"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(string)
    w.Write([]byte("Hello, user " + userID))
}

func RegisterRoutes(router *mux.Router, service *services.Service) {
    router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")

    // Other routes...
	router.HandleFunc("/tasks", getTasks).Methods("GET")
    router.HandleFunc("/tasks", createTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
    router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
}
// // Function to register routes
// func registerRoutes(router *mux.Router) {
//     router.HandleFunc("/tasks", getTasks).Methods("GET")
//     router.HandleFunc("/tasks", createTask).Methods("POST")
//     router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
//     router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
//     router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
// }
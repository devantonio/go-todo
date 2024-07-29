package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "appname/internal/models"
    "appname/internal/services"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type TaskHandler struct {
    taskService *services.TaskService
}

func RegisterRoutes(router *mux.Router, client *mongo.Client) {
    taskService := services.NewTaskService(client)
    handler := &TaskHandler{taskService: taskService}

    router.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
    router.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", handler.GetTask).Methods("GET")
    router.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.taskService.GetTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    task.ID = primitive.NewObjectID()
    _, err := h.taskService.CreateTask(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
    id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    task, err := h.taskService.GetTaskByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
    id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    _, err = h.taskService.UpdateTask(id, task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = h.taskService.DeleteTask(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(string)
    w.Write([]byte("Hello, user " + userID))
}
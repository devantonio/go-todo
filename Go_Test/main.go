package main

import (
	"context"
	"go_test/config"
	"go_test/controllers"
	"go_test/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.DBURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	taskController := controllers.NewTaskController(client, cfg.DBName)

	router := mux.NewRouter()
	routes.RegisterRoutes(router, taskController)

	log.Println("Server starting on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}

package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "appname/internal/config"
    "appname/internal/db"
    "appname/internal/api/handlers"
    "appname/internal/api/middleware"
    "appname/internal/api/handlers"
)

func main() {
    cfg := config.LoadConfig()
    dbClient := db.Connect(cfg.DBURI)

    router := mux.NewRouter()

    // Serve static files
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

    // Use the JWTAuthMiddleware for all routes
    router.Use(middleware.JWTAuthMiddleware)

    // Register API routes
    handlers.RegisterRoutes(router, dbClient)

    // Serve the index.html template
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "web/templates/index.html")
    })

    log.Println("Server starting on port", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}

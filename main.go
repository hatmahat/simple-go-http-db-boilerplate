package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"simple-http-boilerplate/handler"
	"simple-http-boilerplate/repository"
	"simple-http-boilerplate/service"

	_ "github.com/lib/pq"
)

// DB connection
const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "testdb"
)

var db *sql.DB

func main() {
	// Initialize the database connection
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Verify the database connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to the database successfully.")

	// Initialize components
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(service.UserServiceParam{UserRepo: userRepo})
	userHandler := handler.NewUserHandler(userService)

	// Register routes
	http.HandleFunc("/users", userHandler.GetUsersHandler)
	http.HandleFunc("/add-user", userHandler.AddUserHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

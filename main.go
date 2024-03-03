package main

import (
	"log"
	"os"
	"time"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/router"
	"github.com/DevloperAmanSingh/go-blog-api/internal/runner"
	"github.com/joho/godotenv"
)

func main() {
	app := router.SetupRouter()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DATABASE_NAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	db.ConnectDatabase(dbURL, dbName)
	defer db.DisconnectDatabase()
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	for {
		select {
		case <-ticker.C:
			runner.CheckSceheduledPosts()
		}
	}

}

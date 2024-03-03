package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var dbName string
var postCollection *mongo.Collection
var userCollection *mongo.Collection
var SchedulePostCollection *mongo.Collection

func ConnectDatabase(url, name string) {
	dbName = name
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ping the database to check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	postCollection = client.Database(dbName).Collection("posts")
	userCollection = client.Database(dbName).Collection("users")
	SchedulePostCollection = client.Database(dbName).Collection("scheduleposts")

	log.Println("Connected to MongoDB!")

}

func GetPostCollection() *mongo.Collection {
	if postCollection == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return postCollection
}

func GetUserCollection() *mongo.Collection {
	if userCollection == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return userCollection
}

func GetSchedulePostCollection() *mongo.Collection {
	if SchedulePostCollection == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return SchedulePostCollection
}

func DisconnectDatabase() {
	if client != nil {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from database: %v", err)
		}
	}
}

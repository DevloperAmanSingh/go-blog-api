package runner

import (
	"context"
	"log"
	"time"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

var scehdulePostTransfer struct {
	ID       string `json:"id" bson:"_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Username string `json:"username"`
}

func CheckSceheduledPosts() {
	currentTimeStr := time.Now().UTC().Format("2006-01-02T15:04")

	filteredPost := bson.M{
		"scheduledat": bson.M{
			"$lte": currentTimeStr,
		},
	}

	cursor, err := db.GetSchedulePostCollection().Find(context.Background(), filteredPost)
	if err != nil {
		log.Printf("Error finding posts: %v\n", err)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var post models.ScehdulePostTransfer
		if err := cursor.Decode(&post); err != nil {
			log.Printf("Error decoding post: %v\n", err)
			continue
		}

		// insert the post into the post collection
		_, err := db.GetPostCollection().InsertOne(context.Background(), post)
		if err != nil {
			log.Printf("Error inserting post: %v\n", err)
			continue
		}

		// remove from the scheduled post collection
		_, err = db.GetSchedulePostCollection().DeleteOne(context.Background(), bson.M{"_id": post.ID})
		if err != nil {
			log.Printf("Error deleting post: %v\n", err)
			continue
		}

		log.Printf("Post scheduled at  inserted into the post collection\n")

	}

}

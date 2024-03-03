package handlers

import (
	"context"
	"fmt"
	"log"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Body     string             `json:"body"`
	username primitive.ObjectID `json:"userId"`
}

func GetPostsByUser(c *fiber.Ctx) error {
	username := c.Query("username")
	collection := db.GetPostCollection()

	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find user",
		})
	}

	// Find all posts associated with the user ID
	cursor, err := collection.Find(context.Background(), bson.M{"username": username})
	if err != nil {
		log.Printf("Error finding posts: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find posts",
		})
	}
	defer cursor.Close(context.Background())
	var posts []Posts
	for cursor.Next(context.Background()) {
		var post Posts
		err := cursor.Decode(&post)
		if err != nil {
			log.Printf("Error decoding post: %v", err)
			continue
		}
		fmt.Println(post)
		posts = append(posts, post)
	}

	return c.JSON(fiber.Map{
		"message": "Posts found successfully",
		"posts":   posts,
	})
}

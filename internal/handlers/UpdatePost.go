package handlers

import (
	"context"
	"log"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var posts models.Posts

func UpdatePost(c *fiber.Ctx) error {
	// get the post id from the request
	postId := c.Query("id")
	var err error
	// Parse request body into a Post struct
	post := new(models.Posts)
	if err := c.BodyParser(post); err != nil {
		return err
	}
	collection := db.GetPostCollection()

	// now update the post in the database
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": postId}, bson.M{"$set": post})
	if err != nil {
		log.Printf("Error updating the post : %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update the post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Post updated successfully",
		"post":    post,
	})
}

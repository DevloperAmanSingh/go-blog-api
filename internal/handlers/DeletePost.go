package handlers

import (
	"context"
	"log"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePostById(c *fiber.Ctx) error {
	var err error
	postId := c.Query("id")
	collection := db.GetPostCollection()
	result := collection.FindOne(context.Background(), bson.M{"_id": postId})
	if result.Err() != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find the post",
		})
	}

	// now delete the post from the database

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": postId})
	if err != nil {
		log.Printf("Error deleting the post : %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete the post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}

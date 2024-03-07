package handlers

import (
	"context"
	"log"
	"strings"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var post models.Posts

func AddPost(c *fiber.Ctx) error {
	post := new(models.Posts)
	if err := c.BodyParser(post); err != nil {
		return err
	}
	if (post.Title == "") || (post.Body == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title and Body are required",
		})
	}

	post.ID = strings.Split(primitive.NewObjectID().String(), "\"")[1]
	collection := db.GetPostCollection()
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		log.Printf("Error inserting post: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add post",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Post added successfully",
		"post":    post,
	})
}

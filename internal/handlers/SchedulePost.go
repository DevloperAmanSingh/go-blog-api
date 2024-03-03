package handlers

import (
	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SchedulePost(c *fiber.Ctx) error {
	collection := db.GetSchedulePostCollection()
	post := new(models.SchedulePost)
	if err := c.BodyParser(post); err != nil {
		return err
	}
	post.ID = primitive.NewObjectID().Hex()
	_, err := collection.InsertOne(c.Context(), post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to schedule post",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Post scheduled successfully",
		"post":    post,
	})

}

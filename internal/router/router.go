package router

import (
	// "go-blog-api/internal/handlers"
	"github.com/DevloperAmanSingh/go-blog-api/internal/controllers"
	"github.com/DevloperAmanSingh/go-blog-api/internal/handlers"
	"github.com/DevloperAmanSingh/go-blog-api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.Home(c)
	})
	app.Post("/add-post", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return handlers.AddPost(c)
	})
	app.Get("/posts", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return handlers.GetPostsByUser(c)
	})

	app.Delete("/delete-post", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return handlers.DeletePostById(c)
	})

	app.Put("/update-post", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return handlers.UpdatePost(c)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.SignUp(c)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c)
	})

	app.Post("/schedule-post", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return handlers.SchedulePost(c)
	})

	return app
}

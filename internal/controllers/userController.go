package controllers

import (
	"context"
	"log"
	"os"
	"time"

	db "github.com/DevloperAmanSingh/go-blog-api/internal/database"
	"github.com/DevloperAmanSingh/go-blog-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	collection := db.GetUserCollection()

	// Check if the user already exists with the same email or username
	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{"$or": []bson.M{
		{"email": user.Email},
		{"username": user.Username},
	}}).Decode(&existingUser)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User with this email or username already exists",
		})
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}
	user.Password = string(hashedPassword)

	// Insert the new user into the database
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	jwt_secret := os.Getenv("JWT_SECRET")
	loginuser := new(models.LoginRequest)
	if err := c.BodyParser(loginuser); err != nil {
		return err
	}

	collection := db.GetUserCollection()
	var users models.User
	err := collection.FindOne(context.Background(), bson.M{"username": loginuser.Username}).Decode(&users)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(loginuser.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = loginuser.Username
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
		Secure:   false,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   signedToken,
	})

}

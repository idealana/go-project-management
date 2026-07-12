package routes

import (
	"log"

	"github.com/idealana/go-project-management/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
}

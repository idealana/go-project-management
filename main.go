package main

import (
	"log"

	"github.com/idealana/go-project-management/config"
	"github.com/idealana/go-project-management/controllers"
	"github.com/idealana/go-project-management/database/seeds"
	"github.com/idealana/go-project-management/repositories"
	"github.com/idealana/go-project-management/routes"
	"github.com/idealana/go-project-management/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seeds.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(
		userRepo,
	)
	userController := controllers.NewUserController(
		userService,
	)

	routes.Setup(app,
		userController,
	)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)

	log.Fatal(app.Listen(":" + port))
}

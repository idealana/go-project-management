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

	boardRepo := repositories.NewBoardRepository()
	boardMemberRepo := repositories.NewBoardMemberRepository()
	boardService := services.NewBoardService(
		boardRepo,
		userRepo,
		boardMemberRepo,
	)
	boardController := controllers.NewBoardController(
		boardService,
	)

	routes.Setup(app,
		userController,
		boardController,
	)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)

	log.Fatal(app.Listen(":" + port))
}

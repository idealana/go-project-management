package routes

import (
	"log"

	"github.com/idealana/go-project-management/config"
	"github.com/idealana/go-project-management/controllers"
	"github.com/idealana/go-project-management/utils"

	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App,
	uc *controllers.UserController,
	bc *controllers.BoardController,
) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	api := app.Group("/api/v1", jwtWare.New(jwtWare.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func (ctx *fiber.Ctx, err error) error {
			return utils.Unauthorized(ctx, "Unathorized", err.Error())
		},
	}))

	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination)
	userGroup.Get("/:id", uc.GetUser)
	userGroup.Put("/:id", uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)

	boardGroup := api.Group("/boards")
	boardGroup.Post("/", bc.CreateBoard)
	boardGroup.Post("/:id/members", bc.AddBoardMembers)
	boardGroup.Put("/:id", bc.UpdateBoard)
	boardGroup.Delete("/:id/members", bc.RemoveBoardMembers)
}

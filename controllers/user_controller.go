package controllers

import (
	"github.com/idealana/go-project-management/models"
	"github.com/idealana/go-project-management/services"
	"github.com/idealana/go-project-management/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(
	s services.UserService,
) *UserController {
	return &UserController{
		service: s,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Failed to Parsing Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Register Failed", err.Error())
	}

	var userResponse models.UserResponse
	copier.Copy(&userResponse, &user)
	
	return utils.Success(ctx, "Register Success", userResponse)
}

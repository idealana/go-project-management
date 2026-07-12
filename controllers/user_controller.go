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

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Failed to Parsing Data", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Invalid Email or Password", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResponse models.UserResponse
	copier.Copy(&userResponse, &user)
	
	return utils.Success(ctx, "Login Successful", fiber.Map{
		"access_token": token,
		"refresh_token": refreshToken,
		"user": userResponse,
	})
}

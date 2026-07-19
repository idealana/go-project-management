package controllers

import (
	"github.com/idealana/go-project-management/models"
	"github.com/idealana/go-project-management/services"
	"github.com/idealana/go-project-management/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{
		service: s,
	}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var err error

	board := new(models.Board)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(&board); err != nil {
		return utils.BadRequest(ctx, "Failed to Parsing Data", err.Error())
	}

	userID, err = uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.BadRequest(ctx, "Invalid Public ID", err.Error())
	}

	board.OwnerPublicID = userID

	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Failed to Create Board", err.Error())
	}

	return utils.Created(ctx, "Board Successfully Created", board)
}

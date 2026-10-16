package controllers

import(
	"time"
	
	"github.com/idealana/go-project-management/models"
	"github.com/idealana/go-project-management/services"
	"github.com/idealana/go-project-management/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CardController struct {
	service services.CardService
}

func NewCardController(s services.CardService) *CardController {
	return &CardController{service: s}
}

func (c *CardController) CreateCard(ctx *fiber.Ctx) error {
	type CreateCardRequest struct {
		ListPublicID string `json:"list_id"`
		Title string `json:"title"`
		Description string `json:"description"`
		DueDate time.Time `json:"due_date"`
		Position int `json:"position"`
	}

	var req CreateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "Invalid Request Body", err.Error())
	}

	card := &models.Card{
		Title: req.Title,
		Description: req.Description,
		DueDate: &req.DueDate,
		Position: req.Position,
	}

	if err := c.service.Create(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "Failed to create card", err.Error())
	}

	return utils.Success(ctx, "Card Successfully created", card)
}

func (c *CardController) UpdateCard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	type updateCardRequest struct {
		ListPublicID string `json:"list_id"`
		Title string `json:"title"`
		Description string `json:"description"`
		DueDate time.Time `json:"due_date"`
		Position int `json:"position"`
	}

	var req updateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "Invalid Request Body", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid Card ID", err.Error())
	}

	card := &models.Card{
		Title: req.Title,
		Description: req.Description,
		DueDate: &req.DueDate,
		Position: req.Position,
		PublicID: uuid.MustParse(publicID),
	}

	if err := c.service.Update(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "Failed to update card", err.Error())
	}

	return utils.Success(ctx, "Card Successfully updated", card)
}

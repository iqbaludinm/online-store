package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateCategory(ctx *fiber.Ctx) error {
	category := new(models.InsertCategory)
	err := ctx.BodyParser(category)
	if err != nil {
		log.Fatalln(err)
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}

	cat, err := h.app.CreateCategory(*category)
	if err != nil {
		return helpers.InternalServerError(ctx, "Failed to add new category", err)
	}
	return helpers.Created(ctx, "Successfully add new category!", cat)
}

func (h HttpServer) GetCategories(ctx *fiber.Ctx) error {
	listCategories, err := h.app.GetCategories()
	if err != nil {
		return helpers.InternalServerError(ctx, "failed to retrieve all categories", err.Error())
	}
	return helpers.Success(ctx, "Succesfully Retrieve All Categories", listCategories)
}

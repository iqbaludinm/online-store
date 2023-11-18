package handlers

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateCategory(ctx *fiber.Ctx) error {
	category := new(models.InsertCategory)
	err := ctx.BodyParser(category)
	if err != nil {
		log.Println(err)
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			res := make([]helpers.Form, len(verr))
			for i, fe := range verr {
				res[i] = helpers.Form{
					Field:   fe.Field(),
					Message: helpers.FormValidationError(fe),
				}
			}
			log.Println(verr)
			return helpers.BadRequest(ctx, "Failed to add new category", res)
		}
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}

	cat, err := h.app.CreateCategory(*category)
	if err != nil {
		return helpers.InternalServerError(ctx, "Failed to add new cateegory", err.Error())
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

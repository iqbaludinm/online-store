package handlers

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateCartProduct(ctx *fiber.Ctx) error {
	cartProduct := new(models.InsertCartProduct)
	err := ctx.BodyParser(cartProduct)
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
			return helpers.BadRequest(ctx, "Failed to add new cart product", res)
		}
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}

	cat, err := h.app.CreateCartProduct(*cartProduct)
	if err != nil {
		return helpers.InternalServerError(ctx, "Failed to add new cart product", err.Error())
	}
	return helpers.Created(ctx, "Successfully add new cart product!", cat)
}

func (h HttpServer) GetCartProducts(ctx *fiber.Ctx) error {
	listCartProducts, err := h.app.GetProducts()
	if err != nil {
		return helpers.InternalServerError(ctx, "failed to retrieve all products", err.Error())
	}
	return helpers.Success(ctx, "Succesfully Retrieve All Products", listCartProducts)
}
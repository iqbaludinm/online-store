package handlers

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateCart(ctx *fiber.Ctx) error {
	cart := new(models.InsertCart)
	err := ctx.BodyParser(cart)
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
			return helpers.BadRequest(ctx, "Failed to add to cart", res)
		}
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}

	cat, err := h.app.CreateCart(*cart)
	if err != nil {
		return helpers.InternalServerError(ctx, "Failed to add new cart", err.Error())
	}
	return helpers.Created(ctx, "Successfully add new caart!", cat)
}

func (h HttpServer) GetCarts(ctx *fiber.Ctx) error {
	listCarts, err := h.app.GetCarts()
	if err != nil {
		return helpers.InternalServerError(ctx, "failed to retrieve all carts", err.Error())
	}
	return helpers.Success(ctx, "Succesfully Retrieve All Carts", listCarts)
}

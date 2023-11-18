package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateCart(ctx *fiber.Ctx) error {
	cart := new(models.InsertCart)
	err := ctx.BodyParser(cart)
	if err != nil {
		log.Println(err)
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

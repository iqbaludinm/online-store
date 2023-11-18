package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) Register(ctx *fiber.Ctx) error {
	user := new(models.RegisterUser)
	err := ctx.BodyParser(user)
	if err != nil {
		log.Println(err)
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}
	
	u, err := h.app.RegisterUser(user)
	if err != nil {
		return helpers.InternalServerError(ctx, err.Error(), "")
	}
	// Create a cart for the registered user
	_, cartErr := h.app.CreateCart(models.InsertCart{UserID: u.ID})
	if cartErr != nil {
			return helpers.InternalServerError(ctx, "Failed to create a cart", cartErr.Error())
	}

	return helpers.Created(ctx, "Register is successfully!", u)
}

func (h HttpServer) Login(ctx *fiber.Ctx) error {
	var user models.LoginUser
	err := ctx.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return err
	}


	u, err := h.app.LoginUser(user)
	if err != nil {
		log.Println(err)
		return helpers.InternalServerError(ctx, err.Error(), "")
	}
	return helpers.Success(ctx, "Login is successfully!", fiber.Map{"token": u})
}
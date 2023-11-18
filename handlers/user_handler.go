package handlers

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) Register(ctx *fiber.Ctx) error {
	user := new(models.RegisterUser)
	err := ctx.BodyParser(user)
	log.Println(err)
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
			return helpers.BadRequest(ctx, "Failed to Add New User", res)
		}
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}
	
	u, err := h.app.RegisterUser(user)
	if err != nil {
		return helpers.InternalServerError(ctx, err.Error(), "")
	}

	return helpers.Created(ctx, "Register is successfully!", u)
}

func (h HttpServer) Login(ctx *fiber.Ctx) error {
	var user models.LoginUser
	err := ctx.BodyParser(&user)
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
			return helpers.BadRequest(ctx, "Please enter valid data!", res)
		}
		return err
	}


	u, err := h.app.LoginUser(user)
	if err != nil {
		log.Println(err)
		return helpers.InternalServerError(ctx, err.Error(), "")
	}
	return helpers.Success(ctx, "Login is successfully!", fiber.Map{"token": u})
}
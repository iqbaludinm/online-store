package handlers


import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateProduct(ctx *fiber.Ctx) error {
	product := new(models.InsertProduct)
	err := ctx.BodyParser(product)
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

	cat, err := h.app.CreateProduct(*product)
	if err != nil {
		return helpers.InternalServerError(ctx, "Failed to add new product", err.Error())
	}
	return helpers.Created(ctx, "Successfully add new product!", cat)
}

func (h HttpServer) GetProducts(ctx *fiber.Ctx) error {
	listProducts, err := h.app.GetProducts()
	if err != nil {
		return helpers.InternalServerError(ctx, "failed to retrieve all products", err.Error())
	}
	return helpers.Success(ctx, "Succesfully Retrieve All Products", listProducts)
}

package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/models"
)

func (h HttpServer) CreateProduct(ctx *fiber.Ctx) error {
	product := new(models.InsertProduct)
	body := ctx.Body()
	err := json.Unmarshal(body, product)
	if err != nil {
		log.Println(err)
		return helpers.InternalServerError(ctx, "Something's wrong with your input", err.Error())
	}

	cat, err := h.app.CreateProduct(*product)
	if err != nil {
		return helpers.InternalServerError(ctx, "Failed to add new product", err.Error())
	}
	return helpers.Created(ctx, "Successfully add new product!", cat)
}

func (h HttpServer) GetProducts(ctx *fiber.Ctx) error {
	category := ctx.Query("category")
	listProductByCategory, err := h.app.GetProducts(category)
	if err != nil {
		return helpers.InternalServerError(ctx, "failed to retrieve all products", err.Error())
	}
	return helpers.Success(ctx, "Succesfully Retrieve All Products", listProductByCategory)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/handlers"
	_ "github.com/iqbaludinm/online-store/middlewares"
)

func RegisterAPIs(app *fiber.App, server handlers.HttpServer) {
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(401)
	})

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	// Register - Login
	v1.Post("/register", server.Register)
	v1.Post("/login", server.Login)

	// Categories
	category := v1.Group("/categories")
	category.Post("/", server.CreateCategory)
	category.Get("/", server.GetCategories)

	// Carts
	cart := v1.Group("/carts")
	cart.Post("/", server.CreateCart)
	cart.Get("/", server.GetCarts)

	// Products
	product := v1.Group("/products")
	product.Post("/", server.CreateProduct)
	product.Get("/", server.GetProducts)
	// Get All Product List with Query Param Category -> /products?category=electronic
	// v1.Get("/products")

	// Cart - Products
	cartProd := v1.Group("/cart-products")
	cartProd.Post("/", server.CreateCartProduct)
	cartProd.Get("/", server.GetCartProducts)

	// // Post product to cart
	// v1.Post("/cart")

	// // Get All Product List in Cart
	// v1.Get("/carts")

	// // Delete Product List in Cart with carts id and product id
	// v1.Delete("/carts/:cartId")

	// // Checkout cart and payment

	// // Swagger
	// v1.Get("/swagger")
}

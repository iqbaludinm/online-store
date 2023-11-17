package routes

import "github.com/gofiber/fiber/v2"

func RegisterAPIs() {
	app := fiber.New()

	apiv1 := app.Group("/api/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	// Get All Product List with Query Param Category -> /products?category=electronic
	apiv1.Get("/products")

	// Post product to cart
	apiv1.Post("/cart")

	// Get All Product List in Cart
	apiv1.Get("/carts")

	// Delete Product List in Cart with carts id and product id
	apiv1.Delete("/carts/:cartId")

	// Register
	apiv1.Post("/register")

	// Login
	apiv1.Post("/login")

	// Checkout cart and payment

	// Swagger 
	apiv1.Get("/swagger")
}

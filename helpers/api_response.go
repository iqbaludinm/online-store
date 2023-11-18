package helpers

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error Response
func BadRequest(ctx *fiber.Ctx, message string, data ...interface{}) error {
	obj := Response{
		Status:  fiber.StatusBadRequest,
		Message: message,
	}

	if len(data) > 0 {
		obj.Data = data[0]
	}
	return ctx.Status(400).JSON(obj)
}

func NotFound(ctx *fiber.Ctx, message string) error {
	obj := Response{
		Status:  fiber.StatusNotFound,
		Message: message,
		Data:    nil,
	}
	return ctx.Status(404).JSON(obj)
}

func InternalServerError(ctx *fiber.Ctx, message string, data interface{}) error {

	obj := Response{
		Status:  fiber.StatusInternalServerError,
		Message: message,
		Data: data,
	}
	return ctx.Status(500).JSON(obj)
}

func Unauthorized(ctx *fiber.Ctx, message string) error {
	obj := Response{
		Status:  fiber.StatusUnauthorized,
		Message: message,
		Data:    ctx.SendStatus(401),
	}
	return ctx.Status(401).JSON(obj)
}

// Success Response
func Created(ctx *fiber.Ctx, message string, data interface{}) error {
	obj := Response{
		Status:  fiber.StatusCreated,
		Message: message,
		Data:    data,
	}

	return ctx.Status(201).JSON(obj)
}

func Success(ctx *fiber.Ctx, message string, data interface{}) error {
	obj := Response{
		Status:  fiber.StatusOK,
		Message: message,
		Data:    data,
	}
	return ctx.Status(200).JSON(obj)
}

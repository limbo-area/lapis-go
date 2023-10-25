package middleware

import (
	"errors"
	"lapis-go/src/exception"

	"github.com/gofiber/fiber/v2"
)

var InternalErrorHandler = func(c *fiber.Ctx, err error) error {
	var e *exception.HttpException
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if errors.As(err, &e) {
		return c.Status(e.StatusCode).JSON(fiber.Map{
			"message": e.Message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
		"path":    c.Path(),
	})
}

var NotFoundHandler = func(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Route " + c.Path() + " Not Found",
	})
}

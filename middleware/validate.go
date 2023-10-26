package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct[T any](payload T) []string {
	errors := make([]string, 0)
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf(
				"%s must be a %s",
				err.Field(),
				err.Tag(),
			))
		}
	}
	return errors
}

func ValidateSchema(schema interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload = schema

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		errors := ValidateStruct(payload)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message":    "Bad Request",
				"statusCode": 400,
				"errors":     errors,
			})
		}

		c.Locals("input", payload)
		return c.Next()
	}
}

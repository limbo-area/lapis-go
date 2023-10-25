package security

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
)

type AuthType string

const (
	Credentials          = "credentials"
	BearerAuth  AuthType = "BearerAuth"
	ApiKeyAuth  AuthType = "ApiKeyAuth"
)

type ISecurity interface {
	Authorize(c *fiber.Ctx) error
	Callback(c *fiber.Ctx, credentials interface{})
	Provider() AuthType
	Scheme() *openapi3.SecurityScheme
}

type Security struct {
	ISecurity
}

func (s *Security) Callback(c *fiber.Ctx, credentials interface{}) {
	c.Locals(Credentials, credentials)
}

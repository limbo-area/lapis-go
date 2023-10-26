package middleware

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/limbo-tree/lapis-go/common/auth"
	"github.com/limbo-tree/lapis-go/common/cache"
	"github.com/limbo-tree/lapis-go/config"
)

func ValidateToken(c *fiber.Ctx) error {
	var accessToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		accessToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		accessToken = c.Cookies("access_token")
	}

	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Need to be logged in",
		})
	}

	publicKey := config.GetConfig().PublicKey
	tokenClaims, err := auth.VerifyToken(accessToken, publicKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	ctx := context.TODO()
	user, err := cache.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "Token is invalid or session has expired",
		})
	}

	c.Locals("user", user)
	return c.Next()
}

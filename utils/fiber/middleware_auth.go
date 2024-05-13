package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/utils/consts"
)

const (
	bearerToken    = "Bearer "
	bearerLength   = len(bearerToken)
	minTokenLength = 10
)

func GetAuthToken(ctx *fiber.Ctx) string {
	tokenString := ctx.Get(consts.AuthHeader)
	if len(tokenString) < minTokenLength {
		return ""
	}
	// If start with "Bearer " then extract the token part
	if tokenString[0:bearerLength] == bearerToken {
		tokenString = tokenString[bearerLength:]
	}
	return tokenString
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	tokenString := GetAuthToken(ctx)
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	claims, err := token.VerifyAccessToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	// Attach claims to context local so it can be accessed in the route handler
	ctx.Locals("claims", claims)
	return ctx.Next()
}

package auth

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/utils/errors"
	. "github.com/qnify/api-server/utils/fiber"
)

func (m *authModule) logout(ctx *fiber.Ctx) error {
	claims := ctx.Locals("claims").(token.TokenData)

	if err := m.redis.Del(ctx.Context(), strconv.Itoa(claims.UserId)).Err(); err != nil {
		return errors.InternalError("error occured while deleting key", err)
	}

	return SendOk(ctx)
}

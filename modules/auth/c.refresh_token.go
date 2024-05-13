package auth

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/utils/errors"
	. "github.com/qnify/api-server/utils/fiber"
	"github.com/redis/go-redis/v9"
)

type refreshRequest struct {
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (m *authModule) refreshToken(ctx *fiber.Ctx) error {
	req := refreshRequest{}
	if err := ParseBody(&req, ctx); err != nil {
		return err
	}

	tokenData, err := token.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		return errors.BadRequest("invalid refresh token")
	}

	savedToken, err := m.redis.Get(ctx.Context(), strconv.Itoa(tokenData.UserId)).Result()
	if err != nil && err != redis.Nil {
		return errors.InternalError("eror while getting token", err)
	} else if req.RefreshToken != savedToken || err == redis.Nil {
		return errors.BadRequest("invalid refresh token")
	}

	return m.sendAndStoreToken(ctx, tokenData)
}

func (m *authModule) sendAndStoreToken(ctx *fiber.Ctx, claim token.TokenData) error {
	accessToken, refreshToken, err := token.GetTokens(claim)
	if err != nil {
		return errors.InternalError("error while generating access token", err)
	}

	if err := m.redis.Set(ctx.Context(), strconv.Itoa(claim.UserId), refreshToken, 0).Err(); err != nil {
		m.log.Warn("error while storing refresh token")
	}

	return SendResponse(ctx, &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

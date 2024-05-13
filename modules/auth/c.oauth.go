package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/modules/auth/verification"
	"github.com/qnify/api-server/utils/errors"
	. "github.com/qnify/api-server/utils/fiber"
)

type oAuthReq struct {
	Provider int64  `json:"provider,omitempty"`
	Code     string `json:"code,omitempty"`
}

func (m *authModule) oAuth(ctx *fiber.Ctx) error {
	req := oAuthReq{}
	if err := ParseBody(&req, ctx); err != nil {
		return err
	}

	var email string

	switch req.Provider {

	case Google:
		token, err := verification.GetGoogleOauthToken(req.Code, m.config.Google)
		if err != nil {
			return SendString(ctx, err.Error())
		}
		userInfo, err := verification.GetGoogleUser(token.AccessToken)
		if err != nil {
			return SendString(ctx, err.Error())
		}
		email = userInfo.Email

	case Apple:
		return errors.BadRequest("provider yet to support")

	default:
		return SendString(ctx, "Unknown provider")

	}

	userId, _, err := m.findUser("", email)
	if err != nil {
		return err
	}

	return m.sendAndStoreToken(ctx, token.TokenData{UserId: userId})
}

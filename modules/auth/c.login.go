package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/utils/errors"
	. "github.com/qnify/api-server/utils/fiber"
)

type loginRequest struct {
	Phone     string   `json:"phone,omitempty"`
	Email     string   `json:"email,omitempty"`
	Otp       string   `json:"otp,omitempty"`
	AuthCodes []string `json:"authCodes,omitempty"`
}

func (m *authModule) login(ctx *fiber.Ctx) error {
	req := loginRequest{}
	if err := ParseBody(&req, ctx); err != nil {
		return err
	}

	userId, phoneOrEmail, err := m.findUser(req.Phone, req.Email)
	if err != nil {
		return err
	}

	verified := false
	for _, authCode := range req.AuthCodes {
		verified = token.VerifyOtpToken(phoneOrEmail, req.Otp, authCode)
		if verified {
			break
		}
	}
	if !verified {
		return errors.Unauthorised("invalid otp")
	}

	return m.sendAndStoreToken(ctx, token.TokenData{UserId: userId})
}

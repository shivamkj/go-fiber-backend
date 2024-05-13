package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/modules/auth/verification"
	"github.com/qnify/api-server/utils/consts"
	"github.com/qnify/api-server/utils/errors"
	. "github.com/qnify/api-server/utils/fiber"
)

type verifyRequest struct {
	Phone     string     `json:"phone,omitempty"`
	Email     string     `json:"email,omitempty"`
	RoboCheck *roboCheck `json:"roboCheck,omitempty"`
}

type roboCheck struct {
	Type int32  `json:"type,omitempty"`
	Code string `json:"code,omitempty"`
}

func (m *authModule) verify(ctx *fiber.Ctx) error {
	req := verifyRequest{}
	if err := ParseBody(&req, ctx); err != nil {
		return err
	}

	// Check if user exists
	if _, _, err := m.findUser(req.Phone, req.Email); err != nil {
		return err
	}

	otp, err := token.GenerateOtp()
	if err != nil {
		return errors.InternalError("error occured while generating otp", err)
	}

	var otpHash string
	if consts.Dev {
		m.log.Debug("otp sent", "otp", otp)
	}

	if req.Email != "" {
		otpHash = token.GetOtpToken(req.Email, otp)
		if err := verification.SendSms(req.Email, otp); err != nil {
			return SendResponse(ctx, "error occured while sending otp")
		}
	} else {
		otpHash = token.GetOtpToken(req.Phone, otp)
		if err := verification.SendMail(req.Email, otp); err != nil {
			return SendResponse(ctx, "error occured while sending otp")
		}
	}

	// handle seurity with redis
	if err := m.redis.Set(ctx.Context(), "foo", ctx.Query("val"), 0).Err(); err != nil {
		panic(err)
	}

	return SendString(ctx, otpHash)
}

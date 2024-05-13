package auth

import (
	"database/sql"

	gofiber "github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/modules/auth/token"
	"github.com/qnify/api-server/modules/auth/verification"
	"github.com/qnify/api-server/utils/fiber"
	"github.com/redis/go-redis/v9"
	"github.com/zerodha/logf"
)

type (
	authModule struct {
		redis  redis.Cmdable
		db     *sql.DB
		log    *logf.Logger
		config AuthConfig
	}

	AuthConfig struct {
		Origin string                        `yaml:"origin"`
		Token  token.TokenConfig             `yaml:"token"`
		Google verification.GoogleAuthConfig `yaml:"google"`
	}
)

func Routes(app *gofiber.App, redis redis.Cmdable, db *sql.DB, logger *logf.Logger, config AuthConfig) {
	controller := authModule{
		redis:  redis,
		db:     db,
		log:    logger,
		config: config,
	}

	// Starts users verification either by mobile or email auth method (other methods not supported).
	// This will send OTP to their email/mobile which can be used by login API to complete auhentication
	// It will send a code that needs to sent to login API along with otp making it stateless
	// Note: user should already exist in the system
	app.Post("/auth/v1/verify", controller.verify)

	// Let user login via email/mobile, takes otp & list of codes received from verify API
	// If OTP is correct (i.e, matches with the code) and user exits in the system
	// then it returns the access token and refresh token
	app.Post("/auth/v1/login", controller.login)

	// OAuth, for login with Google and Apple
	app.Post("/auth/v1/oauth", controller.oAuth)

	// Refresh Access Token
	app.Post("/auth/v1/refreshToken", controller.refreshToken)

	// Logout User
	app.Delete("/auth/v1/logout", fiber.AuthMiddleware, controller.logout)
}

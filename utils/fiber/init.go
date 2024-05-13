package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/consts"
	"github.com/zerodha/logf"
)

const serverReadTimeout = time.Second * 30

var (
	logger *logf.Logger
)

func InitFiber(log *logf.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  serverReadTimeout,
		ErrorHandler: handleError,
	})

	app.Use(handlePanic)

	if consts.Dev {
		app.Static("/", "./public")
	}

	logger = log

	return app
}

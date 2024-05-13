package fiber

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/consts"
	"github.com/qnify/api-server/utils/errors"
)

func handleError(ctx *fiber.Ctx, err error) error {
	switch t := err.(type) {

	case *fiber.Error:
		if t.Code == fiber.StatusMethodNotAllowed || t.Code == fiber.StatusNotFound {
			return ctx.Status(fiber.StatusNotFound).SendString("Not Found")
		}
		logger.Error("unknown fiber error", "code", t.Code, "err", t.Message)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")

	case *errors.HttpError:
		if consts.Dev {
			logger.Debug("http error", "err", t.ErrorMsg, "code", t.Code)
		}
		return ctx.Status(int(t.Code)).SendString(t.ErrorMsg)

	case *errors.InternalHttpError:
		if consts.Dev {
			fmt.Println(t.Stack())
		}
		logger.Error("internal server error", "err", t.ErrorMsg, "internalError", t.Error(), "stack", t.Stack())
		return ctx.Status(int(t.Code)).SendString("Intenal Server Error")

	case *errors.Err:
		logger.Error("unhandeled error", "err", t.Error(), "stack", t.Stack())
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")

	default:
		logger.Error("unknown error", "err", t.Error())
		ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Internal Error",
		})
	}

	return nil
}

func handlePanic(ctx *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if consts.Dev {
				logger.Error("==== panic occured in fiber handler =====", "err", r)
				debug.PrintStack()
			} else {
				logger.Error("panic occured in fiber handler", "err", r, "stack", debug.Stack())
			}

			// set error for global error handler
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	return ctx.Next()
}

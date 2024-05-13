package fiber

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	. "github.com/qnify/api-server/utils/helper"
)

func TestHandlePanic(t *testing.T) {
	errMsg := "panic error"

	app := fiber.New(fiber.Config{ErrorHandler: handleError})
	var b bytes.Buffer
	logger = InitTestLogger(&b)

	app.Use(handlePanic)
	app.Get("/panic", func(ctx *fiber.Ctx) error {
		panic(errMsg)
	})

	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/panic", nil))
	NoErr(err, t)
	Check(strings.Contains(b.String(), errMsg), t)
	Check(resp.StatusCode == http.StatusInternalServerError, t)
}

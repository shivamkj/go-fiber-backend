package fiber

import (
	"encoding/json"

	goFiber "github.com/gofiber/fiber/v2"
	"github.com/qnify/api-server/utils/consts"
	"github.com/qnify/api-server/utils/errors"
)

type FiberCtx goFiber.Ctx

// Parse JSON from request body
func ParseJSON(ctx *goFiber.Ctx, req any) error {
	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		return errors.BadRequest("error parsing request body")
	}
	return nil
}

// Parse Body with support for protobuf format
func ParseBody(req any, ctx *goFiber.Ctx) error {
	err := json.Unmarshal(ctx.Body(), &req)
	if err != nil {
		return errors.BadRequest("invalid request body")
	}
	return nil
}

func SendResponse(ctx *goFiber.Ctx, response any) error {
	body, err := json.Marshal(response)
	ctx.Set(consts.ContentType, consts.JsonType)
	if err != nil {
		return errors.InternalError("error occured while marshalling", err)
	}
	return ctx.Send(body)
}

func SendString(ctx *goFiber.Ctx, text string) error {
	ctx.Status(200)
	ctx.Set(consts.ContentType, consts.TextType)
	return ctx.SendString(text)
}

func SendOk(ctx *goFiber.Ctx) error {
	return SendString(ctx, "OK")
}

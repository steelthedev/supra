package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHtmxRequest() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Get("Hx-Request") == "true" {
			ctx.Locals("IsHtmxRequest", true)
		} else {
			ctx.Locals("IsHtmxRequest", false)
		}
		return ctx.Next()

	}

}

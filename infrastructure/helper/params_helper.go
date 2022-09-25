package helper

import (
	"github.com/gofiber/fiber/v2"
)

func GetParams(ctx *fiber.Ctx) (uint, error) {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

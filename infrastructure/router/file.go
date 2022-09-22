package router

import (
	"github.com/gofiber/fiber/v2"
	"taveler/interface/controller"
)

func setupFileRouter(api fiber.Router, c controller.FileController) {
	place := api.Group("/file")
	place.Post("", c.UploadImage)
	place.Get("/:id", c.GetFileByID)
}

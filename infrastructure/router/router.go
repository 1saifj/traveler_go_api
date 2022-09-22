package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"taveler/interface/controller"
)

func SetupRouter(app *fiber.App, c *controller.AppController) {
	api := app.Group("/api", logger.New())
	setupControllerRouters(api, c)
}

func setupControllerRouters(api fiber.Router, c *controller.AppController) {
	setupFileRouter(api, *c)
	setupPlaceRouter(api, *c)
	setupSwaggerRouter(api, *c)
	setupUserRouter(api, *c)
}

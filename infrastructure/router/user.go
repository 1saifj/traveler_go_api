package router

import (
	"github.com/gofiber/fiber/v2"
	"taveler/interface/controller"
)

func setupUserRouter(api fiber.Router, c controller.UserController) {
	place := api.Group("/user")
	place.Post("/sign-up", c.CreateUser)
}

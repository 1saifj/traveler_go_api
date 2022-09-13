package router

import (
	"github.com/gofiber/fiber/v2"
	"taveler/interface/controller"
)

func setupPlaceRouter(api fiber.Router, c controller.PlaceController) {
	place := api.Group("/place")
	place.Post("/create", c.CreatePlace)
	place.Get("", c.FindAll)
}

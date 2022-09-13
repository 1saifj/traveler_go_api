package router

import (
	"github.com/kataras/iris/v12/core/router"
	"taveler/interface/controller"
)

func setupPlaceRouter(api router.Party, c controller.PlaceController) {
	place := api.Party("/place")
	place.Post("/create", c.CreatePlace)
	place.Get("", c.FindAll)
}

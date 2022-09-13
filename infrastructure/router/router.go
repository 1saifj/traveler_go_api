package router

import (
	"github.com/kataras/iris/v12/core/router"
	"taveler/interface/controller"
)

func SetupRouter(builder *router.APIBuilder, c *controller.AppController) {
	api := builder.Party("/api")

	setupControllerRouters(api, c)
}

func setupControllerRouters(api router.Party, c *controller.AppController) {
	setupPlaceRouter(api, *c)

}

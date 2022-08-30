package router

import (
	"github.com/kataras/iris/v12/core/router"
	"taveler/config"
	"taveler/interface/controller"
)

func SetupRouter(builder *router.APIBuilder, c *controller.AppController, config *config.AppConfig) {
	api := builder.Party("/api")

	// TODO: enable CORS for web
	// TODO: Setup versioning with iris router
	setupControllerRouters(api, c)
}

func setupControllerRouters(api router.Party, c *controller.AppController) {
	setupPlaceRouter(api, *c)

}

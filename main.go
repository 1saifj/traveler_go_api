package main

import (
	"github.com/kataras/iris/v12"
	"os"
	"taveler/infrastructure/datastore"
	"taveler/infrastructure/router"
	"taveler/registry"
)

// @title           Swagger Example API
// @version         1.0
// @description
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

func main() {

	db, err := datastore.SetupDB()
	app := newApp()

	if err != nil {
		panic(err)
		return
	}
	reg := &registry.Registry{
		DB: db,
	}

	router.SetupRouter(app.APIBuilder, reg.NewAppController())

	hostPort := os.Getenv("HOST_PORT")
	listen(app, hostPort)

}
func newApp() *iris.Application {
	app := iris.Default()
	return app
}

func listen(app *iris.Application, host string) {
	err := app.Listen(host)
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
}

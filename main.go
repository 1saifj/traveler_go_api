package main

import (
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()
	if err != nil {
		panic(err)
		return
	}
	reg := &registry.Registry{
		DB: db,
	}
	router.SetupRouter(app, reg.NewAppController())
	//base := os.Getenv("HOST_PORT")
	err = app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}

}

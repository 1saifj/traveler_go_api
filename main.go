package main

import (
	"github.com/gofiber/fiber/v2"
	_ "taveler/docs"
	"taveler/infrastructure/datastore"
	"taveler/infrastructure/router"
	"taveler/registry"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
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

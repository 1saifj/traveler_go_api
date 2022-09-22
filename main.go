package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
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
	app.Static("/public", "./public")
	if err != nil {
		panic(err)
		return
	}
	reg := &registry.Registry{
		DB: db,
	}
	router.SetupRouter(app, reg.NewAppController())

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	if err != nil {
		log.Fatal("Could not initialize JWT Role Authorizer")
	}

	base := os.Getenv("APP_ADDR")
	err = app.Listen(base)
	if err != nil {
		panic(err)
	}
}

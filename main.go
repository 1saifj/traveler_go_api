package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	_ "taveler/docs"
	"taveler/infrastructure/datastore"
	"taveler/infrastructure/model"
	"taveler/infrastructure/router"
	"taveler/infrastructure/utils"
	"taveler/interface/middlewares"
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

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.All("/api/auth/*", middlewares.MustAuth, func(c *fiber.Ctx) error {
		session, err := datastore.SessionStore.Get(c)
		if err != nil {
			return c.Status(500).JSON("Failed to connect to sessions")
		}
		uid := session.Get("uid")
		if uid == nil {
			return c.Status(500).JSON("Not logged in")
		}

		user := new(model.User)
		result := db.Where("id = ?", uid).First(user)
		if result.Error != nil {
			_ = session.Destroy()
			_ = session.Save()
			return c.Status(500).JSON("User not found")
		}

		if user.State != "Active" {
			return c.Status(401).JSON("User state must be active")
		}
		jwtToken, err := utils.GenerateJWT(user)
		if err != nil {
			return c.Status(500).JSON("Failed to generate token")
		}

		jwtToken = "Bearer " + jwtToken
		c.Set("Authorization", jwtToken)

		return c.SendStatus(200)
	})

	base := os.Getenv("APP_ADDR")
	err = app.Listen(base)
	if err != nil {
		panic(err)
	}
}

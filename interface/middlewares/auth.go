package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"taveler/infrastructure/datastore"
	"taveler/infrastructure/model"
)

func MustAuth(c *fiber.Ctx) error {
	path := strings.Replace(c.Path(), "/api/auth", "", 1)
	if strings.Contains(path, "/api/v2/myauth/identity") || strings.Contains(path, "/api/v2/myauth/public") || strings.Contains(path, "api/v2/product/public") {
		return c.SendStatus(200)
	}
	session, err := datastore.SessionStore.Get(c)

	if err != nil {
		return c.Status(500).JSON("Failed to connect to sessions")
	}
	uid := session.Get("uid")
	if uid == nil {
		return c.Status(401).JSON("Not logged in")
	}
	user := new(model.User)
	result := datastore.DB.Where("id = ?", uid).First(user)
	if result.Error != nil {
		session.Destroy()
		session.Save()
		return c.Status(401).JSON("User not found")
	}
	return c.Next()
}

//func MustPending(c *fiber.Ctx) error {
//	user := c.Locals("CurrentUser").(*models.User)
//
//	if user.State != "Pending" {
//		return c.Status(422).JSON("User state must be pending")
//	}
//
//	return c.Next()
//}
//
//func MustGuest(c *fiber.Ctx) error {
//	session, err := config.SessionStore.Get(c)
//
//	if err != nil {
//		return c.Status(500).JSON(controllers.FailedConnectToSessions)
//	}
//
//	uid := session.Get("uid")
//
//	if uid == nil {
//		return c.Next()
//	}
//
//	return c.Status(422).JSON("Must be guest")
//}
//
//func CheckRequest(c *fiber.Ctx) error {
//	jwt_auth, err := utils.CheckJWT(strings.Replace(c.Get("Authorization"), "Bearer ", "", -1))
//
//	if err != nil {
//		return c.Status(500).JSON(controllers.FailedToParseJWT)
//	}
//
//	user := new(models.User)
//
//	if err := collection.User.FindOne(context.Background(), bson.M{"uid": jwt_auth.UID}).Decode(&user); err != nil {
//		return c.Status(500).JSON(controllers.ServerInternalError)
//	}
//
//	if len(user.Email) == 0 {
//		return c.Status(500).JSON(controllers.ServerInternalError)
//	}
//
//	c.Locals("CurrentUser", user)
//
//	return c.Next()
//}

package datastore

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"log"
	"os"
	"time"
)

var SessionStore *session.Store

func InitSessionStore() {
	storage := redis.New(redis.Config{
		Host:  os.Getenv("REDIS_HOST"),
		Port:  6379,
		Reset: false,
	})

	SessionStore = session.New(session.Config{
		Storage:        storage,
		Expiration:     7 * time.Hour,
		CookiePath:     "/",
		CookieHTTPOnly: true,
	})

	log.Println("Initialize session success")
}

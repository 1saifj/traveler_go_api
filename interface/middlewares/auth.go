package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	jwt_utils "taveler/infrastructure/utils"
)

// JWTRoleAuthorizer is a sturcture for a Role Authorizer type
type JWTRoleAuthorizer struct {
	enforcer   *casbin.Enforcer
	SigningKey []byte
	//	logger     logger.Logger
}

const CasbinConfigPath = "./config/rbac_model.conf"
const MiddlewareRolesPath = "./config/models.csv"
const JWTSecretKey = "secret"

func NewJWTRoleAuthorizer(MiddlewareRolesPath interface{}) (*JWTRoleAuthorizer, error) {
	enforcer, err := casbin.NewEnforcer(CasbinConfigPath, MiddlewareRolesPath)
	if err != nil {
		log.Fatal("could not initialize new enforcer:", err.Error())
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		SigningKey: []byte(JWTSecretKey),
		//		logger:     logger,
	}, nil
}

func NewAuthorizer(jwtra *JWTRoleAuthorizer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accessToken := ctx.Get("Authorization")
		claims, err := jwt_utils.ExtractClaims(accessToken, jwtra.SigningKey)
		if err != nil {
			log.Println("could not extract claims:", err)
			return err
		}

		role := claims["role"]
		ok, err := jwtra.enforcer.Enforce(role, ctx.Path(), ctx.Method())
		if err != nil {
			log.Println("could not enforce:", err)
			return err
		}

		if !ok {
			err = ctx.SendStatus(http.StatusForbidden)
			if err != nil {
				return err
			}
			return ctx.JSON(fiber.Map{
				"message": "You are not authorized to access this resource",
			})
		}

		return ctx.Next()
	}
}

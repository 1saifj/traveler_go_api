package controller

import (
	"github.com/gofiber/fiber/v2"
	"taveler/infrastructure/datastore"
	"taveler/infrastructure/model"
	"taveler/infrastructure/parameter"
	"taveler/infrastructure/utils"
	"taveler/usecase/interactor"
)

type userController struct {
	interactor interactor.UserInteractor
}

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
}

func NewUserController(i interactor.UserInteractor) UserController {
	return &userController{interactor: i}
}

func (u *userController) CreateUser(ctx *fiber.Ctx) error {
	body := new(parameter.SignUp)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	user := &model.User{
		Email: body.Email,
	}

	exists, _ := u.interactor.UserExists(user.Email)
	sess, err := datastore.SessionStore.Get(ctx)
	if exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(400,
			"Email already exists", "SignUp"))
	}
	//create user
	user.Password = utils.HashPassword(body.Password)
	createdUser, err := u.interactor.CreateUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	sess.Set("uid", createdUser.ID)
	_ = sess.Save()
	return ctx.JSON(createdUser)
}

package controller

import (
	"github.com/gofiber/fiber/v2"
	"taveler/infrastructure/model"
	"taveler/infrastructure/parameter"
	"taveler/infrastructure/service/authorization"
	"taveler/infrastructure/utils"
	"taveler/usecase/interactor"
)

type userController struct {
	interactor  interactor.UserInteractor
	accessToken authorization.TokenService
}

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

func NewUserController(i interactor.UserInteractor, at authorization.TokenService) UserController {
	return &userController{interactor: i,
		accessToken: at}

}

func (u *userController) CreateUser(ctx *fiber.Ctx) error {
	body := new(parameter.SignUp)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	user := &model.User{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	exists, err := u.interactor.UserExists(user.Email)
	if exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User already exists",
		})
	}
	//create user
	user.Password = utils.HashPassword(body.Password)
	createdUser, err := u.interactor.CreateUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(createdUser)
}

func (u *userController) Login(ctx *fiber.Ctx) error {
	body := new(parameter.LoginIn)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := utils.ValidateStruct(body); err != nil {
		return err
	}

	user, err := u.interactor.FindByEmail(body.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	//check password hash
	isPasswordCorrect := utils.CheckPasswordHash(body.Password, user.Password)
	if !isPasswordCorrect {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(400,
			err, "Incorrect password"))
	}
	atk, err := u.accessToken.GenerateAccessToken(ctx, user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	rtk, err := u.accessToken.GenerateRefreshToken(ctx, user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	us, err := u.interactor.Login(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	return ctx.JSON(fiber.Map{
		"access_token":  atk,
		"refresh_token": rtk,
		"user":          us,
	})

}

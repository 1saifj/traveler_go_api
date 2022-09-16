package controller

import (
	"github.com/gofiber/fiber/v2"
	"taveler/infrastructure/parameter"
	"taveler/usecase/interactor"
)

type placeController struct {
	Interactor interactor.PlaceInteractor
}

type PlaceController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx)
	CreatePlace(ctx *fiber.Ctx) error
	UpdatePlace(ctx *fiber.Ctx)
	DeletePlace(ctx *fiber.Ctx)
}

func NewPlaceController(i interactor.PlaceInteractor) PlaceController {
	return &placeController{Interactor: i}
}

func (p *placeController) FindAll(ctx *fiber.Ctx) error {
	var filter parameter.Filter
	err := ctx.QueryParser(&filter)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(err)
	}
	places, er := p.Interactor.FindAll(filter)
	if er != nil {
		panic(err)
	}
	return ctx.JSON(places)

}

func (p *placeController) FindByID(ctx *fiber.Ctx) {
	//TODO implement me
	panic("implement me")
}

func (p *placeController) CreatePlace(ctx *fiber.Ctx) error {
	param := parameter.Place{}
	err := ctx.BodyParser(&param)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(err)
	}
	place, err := p.Interactor.CreatePlace(param)
	if err != nil {
		ctx.Status(fiber.StatusForbidden).JSON(err)
		panic(err)
	}
	return ctx.JSON(place)
}

func (p *placeController) UpdatePlace(ctx *fiber.Ctx) {
	//TODO implement me
	panic("implement me")
}

func (p *placeController) DeletePlace(ctx *fiber.Ctx) {
	//TODO implement me
	panic("implement me")
}

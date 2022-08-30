package controller

import (
	"github.com/kataras/iris/v12"
	"taveler/infrastructure/parameter"
	"taveler/usecase/interactor"
)

type placeController struct {
	Interactor interactor.PlaceInteractor
}

type PlaceController interface {
	FindAll(ctx iris.Context)
	FindByID(ctx iris.Context)
	CreatePlace(ctx iris.Context)
	UpdatePlace(ctx iris.Context)
	DeletePlace(ctx iris.Context)
}

func NewPlaceController(i interactor.PlaceInteractor) PlaceController {
	return &placeController{Interactor: i}
}

func (p *placeController) FindAll(ctx iris.Context) {
	//TODO implement me
	panic("implement me")
}

func (p *placeController) FindByID(ctx iris.Context) {
	//TODO implement me
	panic("implement me")
}

func (p *placeController) CreatePlace(ctx iris.Context) {
	param := parameter.Place{}
	err := ctx.ReadJSON(&param)
	if err != nil {
		return
	}
	place, err := p.Interactor.CreatePlace(param)
	if err != nil {
		//ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	_, _ = ctx.JSON(place)
}

func (p *placeController) UpdatePlace(ctx iris.Context) {
	//TODO implement me
	panic("implement me")
}

func (p *placeController) DeletePlace(ctx iris.Context) {
	//TODO implement me
	panic("implement me")
}

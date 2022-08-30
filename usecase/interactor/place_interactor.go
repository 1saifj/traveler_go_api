package interactor

import (
	"taveler/infrastructure/model"
	"taveler/infrastructure/parameter"
	"taveler/usecase/presenter"
	"taveler/usecase/repository"
)

type placeInteractor struct {
	repository repository.PlaceRepository
	presenter  presenter.PlacePresenter
}

type PlaceInteractor interface {
	FindAll(parameter.Filter) ([]any, error)
	FindByID(id uint) (*model.Place, error)
	CreatePlace(parameter.Place) (*model.Place, error)
	UpdatePlace(id uint, params parameter.Place) (*model.Place, error)
	DeletePlace(id uint) error
}

func NewPlaceInteractor(r repository.PlaceRepository, p presenter.PlacePresenter) PlaceInteractor {
	return &placeInteractor{repository: r, presenter: p}
}

func (c *placeInteractor) FindAll(filter parameter.Filter) ([]any, error) {
	places, err := c.repository.FindAll(filter)
	if err != nil {
		return nil, err
	}
	return c.presenter.ResponsePlaces(places), nil
}

func (c *placeInteractor) CreatePlace(param parameter.Place) (*model.Place, error) {
	return c.repository.CreatePlace(param)
}

func (c *placeInteractor) FindByID(id uint) (*model.Place, error) {
	return c.repository.FindByID(id)
}

func (c *placeInteractor) UpdatePlace(id uint, params parameter.Place) (*model.Place, error) {
	return c.repository.UpdatePlace(id, params)
}

func (c *placeInteractor) DeletePlace(id uint) error {
	return c.repository.DeletePlace(id)
}

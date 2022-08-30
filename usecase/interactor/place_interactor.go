package interactor

import (
	"taveler/infrastructure/parameter"
	"taveler/usecase/presenter"
	"taveler/usecase/repository"
)

type placeInteractor struct {
	repository repository.PlaceRepository
	presenter  presenter.PlacePresenter
}

type PlaceInteractor interface {
	GetPlaces(filter parameter.Filter) ([]any, error)
}

func NewPlaceInteractor(r repository.PlaceRepository, p presenter.PlacePresenter) PlaceInteractor {
	return &placeInteractor{repository: r, presenter: p}
}

func (c *placeInteractor) GetPlaces(filter parameter.Filter) ([]any, error) {
	places, err := c.repository.FindAll(filter)
	if err != nil {
		return nil, err
	}
	return c.presenter.ResponsePlaces(places)
}

package repository

import (
	"taveler/infrastructure/model"
	"taveler/infrastructure/parameter"
)

type PlaceRepository interface {
	FindAll(parameter.Filter) ([]*model.Place, error)
	FindByID(id uint) (*model.Place, error)
	CreatePlace(parameter.Place) (*model.Place, error)
	UpdatePlace(id uint, params parameter.Place) (*model.Place, error)
	DeletePlace(id uint) error
}

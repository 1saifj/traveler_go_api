package presenter

import (
	"taveler/infrastructure/model"
)

type PlacePresenter interface {
	ResponsePlaces(places []*model.Place) []any
}

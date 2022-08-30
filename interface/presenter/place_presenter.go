package presenter

import "taveler/infrastructure/model"

type placePresenter struct {
}

type PlacePresenter interface {
	ResponsePlaces([]*model.Place) []any
}

func NewPlacePresenter() PlacePresenter {
	return &placePresenter{}
}

func (c *placePresenter) ResponsePlace(place model.Place) any {
	return map[string]any{
		"id":          place.ID,
		"name":        place.Name,
		"address":     place.Address,
		"description": place.Description,
		"latitude":    place.Latitude,
		"longitude":   place.Longitude,
		"slug":        place.Slug,
		"thumbnail":   place.Thumbnail,
		"likes":       place.Likes,
		"comments":    place.Comments,
	}
}

func (c *placePresenter) ResponsePlaces(places []*model.Place) []any {
	data := make([]any, 0)
	for _, place := range places {
		data = append(data, c.ResponsePlace(*place))
	}
	return data
}

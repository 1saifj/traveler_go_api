package repository

import (
	"gorm.io/gorm"
	"taveler/infrastructure/model"
	"taveler/infrastructure/parameter"
)

type placeRepository struct {
	db *gorm.DB
}

type PlaceRepository interface {
	FindAll(parameter.Filter) ([]*model.Place, error)
	FindByID(id uint) (*model.Place, error)
	CreatePlace(parameter.Place) (*model.Place, error)
	UpdatePlace(id uint, params parameter.Place) (*model.Place, error)
	DeletePlace(id uint) error
	SaveImage(place parameter.Image)
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return &placeRepository{db: db}
}

func (p *placeRepository) FindAll(param parameter.Filter) ([]*model.Place, error) {
	places := []*model.Place{}
	err := p.db.Limit(param.GetLimit()).Offset(param.GetOffset()).Order(param.OrderQueryBy()).Find(&places).Error
	if err != nil {
		return nil, err
	}
	return places, nil

}

func (p *placeRepository) FindByID(id uint) (*model.Place, error) {
	//TODO implement me
	panic("implement me")

}

func (p *placeRepository) CreatePlace(params parameter.Place) (*model.Place, error) {
	place := model.Place{
		Name:        params.Name,
		Description: params.Description,
		Latitude:    params.Latitude,
		Longitude:   params.Longitude,
		Address:     params.Address,
		Slug:        params.Slug,
		//Thumbnail:   params.Thumbnail,
	}
	err := p.db.Create(place).Error
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func (p *placeRepository) UpdatePlace(id uint, params parameter.Place) (*model.Place, error) {
	//TODO implement me

	panic("implement me")
}

func (p *placeRepository) DeletePlace(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (p *placeRepository) SaveImage(img parameter.Image) {
	p.db.Save(&img)
	p.db.Preload("Images").Find(&img)
}

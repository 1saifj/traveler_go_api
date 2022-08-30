package registry

import (
	"gorm.io/gorm"
	"taveler/interface/controller"
)

type registry struct {
	DB *gorm.DB
}

type Registry interface {
	NewAppController() *controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() *controller.AppController {
	return &controller.AppController{
		PlaceController: r.NewPlaceController(),
	}
}

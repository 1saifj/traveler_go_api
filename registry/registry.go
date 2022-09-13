package registry

import (
	"gorm.io/gorm"
	"taveler/interface/controller"
)

type Registry struct {
	DB *gorm.DB
}

func (r *Registry) NewAppController() *controller.AppController {
	return &controller.AppController{
		PlaceController: r.NewPlaceController(),
	}
}

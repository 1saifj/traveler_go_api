package registry

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"taveler/interface/controller"
)

type Registry struct {
	DB     *gorm.DB
	Logger zerolog.Logger
}

func (r *Registry) NewAppController() *controller.AppController {
	return &controller.AppController{
		FileController:  r.NewFileController(),
		PlaceController: r.NewPlaceController(),
		UserController:  r.NewUserController(),
	}
}

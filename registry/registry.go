package registry

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"taveler/infrastructure/service/authorization"
	"taveler/interface/controller"
)

type Registry struct {
	DB     *gorm.DB
	Logger zerolog.Logger
	Token  authorization.TokenService
}

func (r *Registry) NewAppController() *controller.AppController {
	return &controller.AppController{
		FileController:  r.NewFileController(),
		PlaceController: r.NewPlaceController(),
		UserController:  r.NewUserController(),
	}
}

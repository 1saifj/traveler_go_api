package registry

import (
	"taveler/interface/controller"
	ip "taveler/interface/presenter"
	ir "taveler/interface/repository"
	"taveler/usecase/interactor"
	up "taveler/usecase/presenter"
	ur "taveler/usecase/repository"
)

func (r *Registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())

}

func (r *Registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *Registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.DB)
}

func (r *Registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

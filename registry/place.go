package registry

import (
	"taveler/interface/controller"
	ip "taveler/interface/presenter"
	ir "taveler/interface/repository"
	"taveler/usecase/interactor"
	up "taveler/usecase/presenter"
	ur "taveler/usecase/repository"
)

func (r *Registry) NewPlaceController() controller.PlaceController {
	return controller.NewPlaceController(r.NewPlaceInteractor())

}

func (r *Registry) NewPlaceInteractor() interactor.PlaceInteractor {
	return interactor.NewPlaceInteractor(r.NewPlaceRepository(), r.NewPlacePresenter())
}

func (r *Registry) NewPlaceRepository() ur.PlaceRepository {
	return ir.NewPlaceRepository(r.DB)
}

func (r *Registry) NewPlacePresenter() up.PlacePresenter {
	return ip.NewPlacePresenter()
}

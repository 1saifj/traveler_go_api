package registry

import (
	"taveler/interface/controller"
	ip "taveler/interface/presenter"
	ir "taveler/interface/repository"
	"taveler/usecase/interactor"
	up "taveler/usecase/presenter"
	ur "taveler/usecase/repository"
)

func (r *registry) NewPlaceController() controller.PlaceController {
	return controller.NewPlaceController(r.NewPlaceInteractor())

}

func (r *registry) NewPlaceInteractor() interactor.PlaceInteractor {
	return interactor.NewPlaceInteractor(r.NewPlaceRepository(), r.NewPlacePresenter())
}

func (r *registry) NewPlaceRepository() ur.PlaceRepository {
	return ir.NewPlaceRepository(r.DB)
}

func (r *registry) NewPlacePresenter() up.PlacePresenter {
	return ip.NewPlacePresenter()
}

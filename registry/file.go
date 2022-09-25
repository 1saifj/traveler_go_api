package registry

import (
	"taveler/interface/controller"
	ip "taveler/interface/presenter"
	ir "taveler/interface/repository"
	"taveler/usecase/interactor"
	up "taveler/usecase/presenter"
	ur "taveler/usecase/repository"
)

func (r *Registry) NewFileController() controller.FileController {
	return controller.NewFileController(r.NewFileInteractor())

}

func (r *Registry) NewFileInteractor() interactor.FileInteractor {
	return interactor.NewFileInteractor(r.NewFileRepository(), r.NewFilePresenter())
}

func (r *Registry) NewFileRepository() ur.FileRepository {
	return ir.NewFileRepository(r.DB)
}

func (r *Registry) NewFilePresenter() up.FilePresenter {
	return ip.NewFilePresenter()
}

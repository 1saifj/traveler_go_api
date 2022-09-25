package interactor

import (
	"taveler/infrastructure/model"
	"taveler/usecase/presenter"
	"taveler/usecase/repository"
)

type fileInteractor struct {
	repository repository.FileRepository
	presenter  presenter.FilePresenter
}

type FileInteractor interface {
	UploadImage(file *model.File) (any, error)
	GetFileByID(id string) (*model.File, error)
}

func NewFileInteractor(r repository.FileRepository, p presenter.FilePresenter) FileInteractor {
	return &fileInteractor{repository: r, presenter: p}
}

func (c *fileInteractor) UploadImage(file *model.File) (any, error) {
	file, err := c.repository.UploadImage(file)
	if err != nil {
		return nil, err
	}
	return c.presenter.UploadImage(file)
}

func (c *fileInteractor) GetFileByID(id string) (*model.File, error) {
	return c.repository.GetFileByID(id)
}

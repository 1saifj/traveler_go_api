package presenter

import "taveler/infrastructure/model"

type FilePresenter interface {
	UploadImage(file *model.File) (any, error)
	GetFileByID(id string) (any, error)
}

package repository

import "taveler/infrastructure/model"

type FileRepository interface {
	UploadImage(file *model.File) (*model.File, error)
	GetFileByID(id string) (*model.File, error)
}

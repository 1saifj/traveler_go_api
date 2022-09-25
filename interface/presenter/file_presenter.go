package presenter

import (
	"taveler/infrastructure/model"
	"taveler/usecase/presenter"
)

type filePresenter struct {
}

type FilePresenter interface {
	UploadImage(file *model.File) (any, error)
	GetFileByID(id string) (any, error)
}

func NewFilePresenter() presenter.FilePresenter {
	return &filePresenter{}
}

func (f *filePresenter) UploadImage(file *model.File) (any, error) {
	return map[string]interface{}{
		"file_id":   file.ID,
		"file_path": file.Path,
	}, nil
}

func (f *filePresenter) GetFileByID(string) (any, error) {
	var file model.File
	return map[string]interface{}{
		"file_path": file.Path,
	}, nil
}

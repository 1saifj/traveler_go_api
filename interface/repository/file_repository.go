package repository

import (
	"gorm.io/gorm"
	"taveler/infrastructure/model"
)

type fileRepository struct {
	DB *gorm.DB
}

type FileRepository interface {
	UploadImage(file *model.File) (*model.File, error)
	GetFileByID(id string) (*model.File, error)
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{DB: db}
}

func (r *fileRepository) UploadImage(file *model.File) (*model.File, error) {
	if err := r.DB.Create(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (r *fileRepository) GetFileByID(id string) (*model.File, error) {
	file := &model.File{}
	if err := r.DB.Where("id = ?", id).First(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

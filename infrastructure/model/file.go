package model

type File struct {
	Model
	FileType string `json:"file_type"`
	FileSize int64  `json:"file_size"`
	Path     string `json:"path" gorm:"type:varchar(255);not null"`
}

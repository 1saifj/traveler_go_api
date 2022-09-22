package model

type File struct {
	Model
	Path string `json:"path" gorm:"type:varchar(255);not null"`
}

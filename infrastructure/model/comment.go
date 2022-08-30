package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"not null"`
}

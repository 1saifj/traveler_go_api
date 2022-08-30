package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string  `json:"name"  gorm:"not null"`
	Places []Place `json:"place"  gorm:"foreignKey:CategoryID"`
}

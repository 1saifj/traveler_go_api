package model

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name        string `json:"name"  gorm:"not null"`
	Description string `json:"description"  gorm:"type:varchar(5000);not null"`
	CategoryID  string `gorm:"type:uuid;not null"  json:"category_id"`
	// Location
	Latitude  float64 `json:"latitude"  form:"latitude" gorm:"type:float;not null"`
	Longitude float64 ` json:"longitude" form:"longitude" gorm:"type:float;not null"`
	Address   string  `json:"address" gorm:"type:varchar(100);not null" `

	Slug      string `json:"slug" gorm:"type:varchar(100);not null"`
	Thumbnail string `json:"thumbnail" gorm:"type:varchar(100); not null"`
	Likes     int    `json:"likes"  gorm:"type:int(10);"`

	Comments []Comment `json:"comments" gorm:"foreignKey:PlaceID"`
}

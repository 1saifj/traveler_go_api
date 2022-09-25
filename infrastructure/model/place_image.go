package model

type Image struct {
	Model
	PlaceID string `json:"place_id" gorm:"not null"`
	Image   string `json:"image" gorm:"type:varchar(255);not null"`
	Place   Place  `gorm:"foreignKey:PlaceID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"place"`
}

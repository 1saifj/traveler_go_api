package model

type Category struct {
	Model
	Name  string `json:"name"  gorm:"not null"`
	Image string `gorm:"type:varchar(255)" json:"image"`
}

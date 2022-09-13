package model

type Place struct {
	Model
	Name        string `json:"name"  gorm:"not null"`
	Description string `json:"description"  gorm:"type:varchar(5000);not null"`
	CategoryID  string `gorm:"not null"  json:"category_id"`
	// Location
	Latitude  float64 `json:"latitude"  form:"latitude" gorm:"type:float;not null"`
	Longitude float64 ` json:"longitude" form:"longitude" gorm:"type:float;not null"`
	Address   string  `json:"address" gorm:"type:varchar(100);not null" `

	Slug     string   `json:"slug" gorm:"type:varchar(100);not null"`
	Likes    int      `json:"likes"  gorm:"type:int(10);"`
	Category Category `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`

	////comments
	//Comments []Comment `json:"comments" gorm:"foreignKey:PlaceID"`
}

type Comment struct {
	Model
	PlaceID uint   `json:"place_id" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	Content string `json:"content" gorm:"type:varchar(5000);not null"`
}

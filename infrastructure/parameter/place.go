package parameter

type Place struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  int    `json:"category_id" form:"category_id" validate:"required"`
	// Location
	Latitude  float64 `json:"latitude" form:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"required"`
	Address   string  `json:"address" form:"address" validate:"required"`
	Slug      string  `json:"slug" form:"slug" validate:"required"`
	Likes     int     `json:"likes" form:"likes" validate:"required"`
}

type Image struct {
	PlaceID string `json:"place_id" form:"place_id" validate:"required"`
	Image   string `json:"image"  form:"image" validate:"required"`
}

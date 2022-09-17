package model

type User struct {
	Model
	FirstName   string `json:"first_name" gorm:"not null"`
	LastName    string `json:"last_name" gorm:"not null"`
	UserName    string `json:"user_name" gorm:"not null"`
	Email       string `json:"email" gorm:"not null;unique"`
	Password    string `json:"-" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	Avatar      string `json:"avatar" gorm:"not null"`
	Birthday    string `json:"birthday" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`
	Country     string `json:"country" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	State       string `json:"state" gorm:"not null"`
	Role        string `json:"role" gorm:"not null"`
}

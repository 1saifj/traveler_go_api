package model

import "time"

type User struct {
	Model
	FirstName          string    `json:"first_name" gorm:"not null"`
	LastName           string    `json:"last_name" gorm:"not null"`
	UserName           string    `json:"user_name" gorm:"not null"`
	Email              string    `json:"email" gorm:"not null;unique"`
	PhoneNumber        string    `json:"phone_number" gorm:"not null"`
	Avatar             string    `json:"avatar" gorm:"not null"`
	Birthday           string    `json:"birthday" gorm:"not null"`
	Gender             string    `json:"gender" gorm:"not null"`
	Country            string    `json:"country" gorm:"not null"`
	Address            string    `json:"address" gorm:"not null"`
	State              string    `json:"state" gorm:"not null"`
	Token              string    `json:"-"`
	Password           string    `json:"-" gorm:"not null"`
	LastLogin          time.Time `json:"last_login,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	LastPasswordChange time.Time `json:"last_password_change,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
}

// Role model

// ChangePassword updates user's password related fields
func (u *User) ChangePassword(hash string) {
	u.Password = hash
	u.LastPasswordChange = time.Now()
}

// UpdateLastLogin updates last login field
func (u *User) UpdateLastLogin(token string) {
	u.Token = token
	u.LastLogin = time.Now()
}

package model

import "time"

var (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	Model
	FirstName          string    `json:"first_name" gorm:"not null" redis:"first_name"`
	LastName           string    `json:"last_name" gorm:"not null" redis:"last_name"`
	UserName           string    `json:"user_name" gorm:"not null" redis:"user_name"`
	Email              string    `json:"email" gorm:"not null;unique" redis:"email"`
	PhoneNumber        string    `json:"phone_number" gorm:"not null" redis:"phone_number"`
	Avatar             string    `json:"avatar" gorm:"not null" redis:"avatar"`
	Birthday           string    `json:"birthday" gorm:"not null" redis:"birthday"`
	Role               string    `json:"role" gorm:"not null" redis:"role"`
	Gender             string    `json:"gender" gorm:"not null"`
	Country            string    `json:"country" gorm:"not null"`
	Address            string    `json:"address" gorm:"not null"`
	State              string    `json:"state" gorm:"not null"`
	Token              string    `json:"-"`
	Password           string    `json:"-" gorm:"not null"`
	LastLogin          time.Time `json:"last_login,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	LastPasswordChange time.Time `json:"last_password_change,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	Tokens             Tokens    `json:"tokens,omitempty" gorm:"-"`
}

type Tokens struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
}

func (u *User) GetUserId() string {
	if u.ID != "" {
		return u.ID
	}
	return "NO_ID"
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsUser() bool {
	return u.Role == RoleUser
}

// ChangePassword updates user's password related fields
func (u *User) ChangePassword(hash string) {
	u.Password = hash
	u.LastPasswordChange = time.Now()
}

//// UpdateLastLogin updates last login field
//func (u *User) UpdateLastLogin(accessToken string) {
//	u.LastLogin = time.Now()
//}

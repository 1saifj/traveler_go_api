package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Model struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	ud := uuid.New().String()
	u.ID = strings.ReplaceAll(ud, "-", "")
	return
}

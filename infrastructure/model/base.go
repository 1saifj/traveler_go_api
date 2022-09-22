package model

import (
	"time"
)

type Model struct {
	ID        uint64    `gorm:"primaryKey;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}

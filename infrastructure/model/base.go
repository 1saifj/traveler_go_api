package model

import (
	"time"
)

type Model struct {
	ID        uint64 `gorm:"primaryKey;auto_increment" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

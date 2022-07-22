package models

import (
	"time"

	"gorm.io/gorm"
)

type Building struct {
	Name      string
	Address   string
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

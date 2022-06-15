package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TokenDetail struct {
	gorm.Model
	UUID       string    `gorm:"not null"`
	Token      string    `gorm:"not null"`
	ExpiryTime time.Time `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	User       User      `gorm:"constraint:onDelete:CASCADE, onUpdate:CASCADE"`
}

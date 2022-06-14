package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type ProductRequestDTO struct {
	Name string `json:"name"`
}

type ProductResponsDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

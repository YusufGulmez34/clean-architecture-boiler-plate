package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Firstname string `gorm:"type:varchar(15);not null"`
	Lastname  string `gorm:"type:varchar(15);not null"`
	Username  string `gorm:"type:varchar(15);not null"`
	Email     string `gorm:"type:varchar(50);not null"`
	Password  string `gorm:"not null"`
}

type UserLoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterDTO struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type UserResponseDTO struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

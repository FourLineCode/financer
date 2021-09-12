package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" validate:"required,min=1,max=30"`
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" gorm:"default=USER" validate:"oneof=USER ADMIN"`
}

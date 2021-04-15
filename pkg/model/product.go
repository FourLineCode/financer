package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=1,max=30"`
	Price    uint   `json:"price" validate:"required"`
	Quantity uint   `json:"quantity" validate:"required"`
}

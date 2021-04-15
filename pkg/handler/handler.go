package handler

import "gorm.io/gorm"

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handler {
	handler := &Handler{
		db: db,
	}

	return handler
}

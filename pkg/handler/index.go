package handler

import (
	"net/http"

	"gorm.io/gorm"
)

type Success struct {
	Success bool `json:"success"`
}

func IndexHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusOK, Success{Success: true})
}

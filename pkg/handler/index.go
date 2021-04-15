package handler

import (
	"net/http"

	"github.com/FourLineCode/financer/pkg/model"
	"gorm.io/gorm"
)

func IndexHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusOK, model.Success{Success: true})
}

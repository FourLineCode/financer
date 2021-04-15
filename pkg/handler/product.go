package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FourLineCode/financer/pkg/model"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	result := []model.Product{}
	db.Find(&result)

	ResponseJSON(w, http.StatusOK, result)
}

func GetProductByID(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := model.Product{}
	if err := db.First(&product, id).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, product)
}

func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validate.Struct(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.(validator.ValidationErrors).Error())
		return
	}

	if err := db.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, product)
}

func UpdateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := model.Product{}
	if err := db.First(&product, id).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, Success{Success: true})
}

func DeleteProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := model.Product{}
	if err := db.First(&product, id).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, Success{Success: true})
}

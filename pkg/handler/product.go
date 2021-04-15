package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FourLineCode/financer/pkg/model"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate = validator.New()

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	result := []model.Product{}
	h.db.Find(&result)

	ResponseJSON(w, http.StatusOK, result)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := model.Product{}
	if err := h.db.First(&product, id).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, product)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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

	if err := h.db.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, product)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := model.Product{}
	if err := h.db.First(&product, id).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, Success{Success: true})
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	product := model.Product{}
	if err := h.db.First(&product, id).Error; err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Delete(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJSON(w, http.StatusOK, Success{Success: true})
}

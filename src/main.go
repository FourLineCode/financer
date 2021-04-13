package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Response struct {
	Ok bool `json:"ok"`
}

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

func main() {
	router := mux.NewRouter()

	db, err := gorm.Open(sqlite.Open("./db/test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{Ok: true})
	})

	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodPost {
			product := Product{Code: "D42", Price: 100}
			db.Create(&product)

			responseJSON(w, http.StatusOK, product)
			return
		}

		if r.Method == http.MethodGet {
			result := []Product{}
			db.Find(&result)

			responseJSON(w, http.StatusOK, result)
			return
		}
	}).Methods(http.MethodGet, http.MethodPost)

	const PORT = ":5000"
	fmt.Println("Server started on http://localhost" + PORT)
	err = http.ListenAndServe(PORT, router)
	if err != nil {
		panic(err)
	}
}

func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

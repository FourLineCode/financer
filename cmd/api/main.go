package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FourLineCode/financer/config"
	"github.com/FourLineCode/financer/pkg/server"
	"github.com/jesseokeya/go-httplogger"
)

func main() {
	config := config.GetConfig()

	server := server.New()
	server.Initialize(config)

	fmt.Println("Server started on http://localhost" + config.Port)
	err := http.ListenAndServe(config.Port, httplogger.Golog(server.Router))
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}

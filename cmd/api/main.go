package main

import (
	"github.com/FourLineCode/financer/config"
	"github.com/FourLineCode/financer/pkg/server"
)

func main() {
	config := config.GetConfig()

	server := server.New(config)

	server.Run(config.Port)
}

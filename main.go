package main

import (
	"github.com/GuilhermeVozniak/go-opportunities/config"
	"github.com/GuilhermeVozniak/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	//initializa config
	err := config.Init()
	if err != nil {
		logger.Errorf("config inicialization error %v",logger)
		return
	}

	// Initialize the routes using gin default settings
	router.InitializeServer()
}
package main

import (
	"github.com/joaogabee/gopportunities/config"
	"github.com/joaogabee/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("MAIN: ")
	err := config.Init()
	if err != nil {
		logger.Err("error: %v", err)
	}
	router.Initializer()
}

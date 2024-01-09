package main

import (
	"github.com/joaogabsoaresf/mygocrm/config"
	"github.com/joaogabsoaresf/mygocrm/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config init error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}

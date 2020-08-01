package server

import (
	"fmt"
	"log"

	"github.com/rohan-luthra/gin-boilerplate/config"
	"github.com/rohan-luthra/gin-boilerplate/logger"
)

func Init() {

	if err := logger.Init(-1); err != nil {
		log.Fatalf("Unable to initalize logger %v", err)
	}

	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("port")))
}

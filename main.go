package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rohan-luthra/gin-boilerplate/config"
	"github.com/rohan-luthra/gin-boilerplate/server"

	_ "github.com/rohan-luthra/gin-boilerplate/docs"
)

// @title template
// @version 1.0
// @description description
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.host.com
// @contact.email support@host.com

// @license.name Apache 2.0

// NewRouter gin server
// @BasePath /
func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}

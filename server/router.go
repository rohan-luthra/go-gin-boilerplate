package server

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/rohan-luthra/gin-boilerplate/logger"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {

	r := gin.New()

	// default logger
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	// custom logger
	r.Use(ginzap.Ginzap(logger.Log, logger.CustomTimeFormat, true))
	r.Use(ginzap.RecoveryWithZap(logger.Log, true))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

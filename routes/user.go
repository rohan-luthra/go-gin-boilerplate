package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohan-luthra/gin-boilerplate/controllers"
)

func SetupUserRoutes(r *gin.Engine) {

	r := r.Group("/users")
	{
		restaurant.GET("/", controllers.GetUsers)

	}

}

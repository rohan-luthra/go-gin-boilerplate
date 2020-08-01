package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary get all users
// @Description get all users
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.User
// @Router /user [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Hello",
	})
}

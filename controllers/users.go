package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct{}

// Retrieve godoc
// @Summary Get user
// @Schemes
// @Description do ping
// @Tags Users
// @Router /users/:id [get]
func (controller UsersController) Retrieve(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"id": claims["id"],
	})
}

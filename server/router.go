package server

import (
	"github.com/dajeo/go-gin-boilerplate/controllers"
	"github.com/dajeo/go-gin-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	usersController := new(controllers.UsersController)
	healthController := new(controllers.HealthController)
	authMiddleware := middlewares.GetJWT()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/authorize", authMiddleware.LoginHandler)
		authGroup.GET("/refresh", authMiddleware.RefreshHandler)
	}

	userGroup := r.Group("usersController")
	{
		userGroup.GET("/:id", usersController.Retrieve)
	}

	r.GET("/healthController", healthController.Status)

	r.GET("/swagger/*any", swagger.WrapHandler(files.Handler))

	return r
}

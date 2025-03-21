package router

import (
	"github.com/gin-gonic/gin"
	"testovoe/internal/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/users")
	{
		api.POST("/", userHandler.CreateUser)
		api.GET("/:id", userHandler.GetUser)
		api.PUT("/:id", userHandler.UpdateUser)
		api.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}

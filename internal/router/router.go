package router

import (
	"go-test/internal/handler" // 记得引入 handler

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", handler.GetUserHandler)
		userGroup.GET("", handler.GetUserHandler)
		userGroup.POST("", handler.HelloHandler)
		userGroup.PUT("/:id", handler.UpdateUserHandler)
		userGroup.DELETE("/:id", handler.DeleteUserHandler)
	}
	return r
}

package router

import (
	"go-test/internal/handler" // 记得引入 handler
	"go-test/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", handler.GetUserHandler)
		userGroup.GET("", handler.GetUserHandler)
		userGroup.POST("", handler.HelloHandler)
		protected := userGroup.Group("")
		protected.Use(middleware.Auth())
		{
			protected.PUT("/:id", handler.UpdateUserHandler)
			protected.DELETE("/:id", handler.DeleteUserHandler)
		}
	}
	return r
}

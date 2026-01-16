package router

import (
	"go-test/internal/handler" // 记得引入 handler
	"go-test/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 访问 http://localhost:8080/swagger/index.html 即可看到文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.Logger())
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", handler.GetUserHandler)
		userGroup.GET("", handler.GetAllUserHandler)
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

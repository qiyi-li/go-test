package main

import (
	"fmt"
	"go-test/internal/router"
	"go-test/internal/store"
	"go-test/internal/config"
	_ "go-test/docs"
)

// @title           Go-Test User API
// @version         1.0
// @description     这是一个基于 Gin + GORM 的用户管理后台 API 文档
// @host            localhost:8080
// @BasePath        /

func main() {
	config.LoadConfig()
	store.Init(config.AppConfig.Database.DSN)
	r := router.SetupRouter()
	addr := ":" + config.AppConfig.Server.Port
	fmt.Printf("Server is running at http://localhost%s\n", addr)
	r.Run(addr)
}

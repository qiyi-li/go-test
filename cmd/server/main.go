package main

import (
	"fmt"
	"go-test/internal/router"
	"go-test/internal/store"
	"go-test/internal/config"
)

func main() {
	config.LoadConfig()
	store.Init(config.AppConfig.Database.DSN)
	r := router.SetupRouter()
	addr := ":" + config.AppConfig.Server.Port
	fmt.Printf("Server is running at http://localhost%s\n", addr)
	r.Run(addr)
}

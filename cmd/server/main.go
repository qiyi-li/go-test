package main

import (
	"fmt"
	"go-test/internal/router"
	"go-test/internal/store"
)

func main() {
	store.Init()
	r := router.SetupRouter()
	fmt.Println("Server is running at http://localhost:8080")
	r.Run(":8080")
}

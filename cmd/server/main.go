package main

import (
	"fmt"
	"go-test/internal/router"
	"go-test/internal/store"
	"net/http"
)

func main() {
	store.Init()
	router.RegisterRoutes()
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

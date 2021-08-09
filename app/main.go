package main

import (
	"demo/package/repository"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type DemoApp struct {
	DB *gorm.DB
}

// Initialize routes and Server
func main() {
	dbConn := repository.NewDbConn()
	demo := DemoApp{
		DB: dbConn,
	}

	handler := demo.routes()

	fmt.Println("Starting server on %s:%s \n", "localhost", "8090")
	err := http.ListenAndServe(":8090", handler)
	if err != nil {
		fmt.Println("Error starting server %s", err)
	}
}

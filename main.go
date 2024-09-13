package main

import (
	"musicservice/src/api/router"
)

func main() {
	router := router.SetupRoutes()

	router.Run(":8080")
}

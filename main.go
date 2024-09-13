package main

import (
	"fmt"
	"musicservice/src/api/router"
	"os"
)

func main() {
	router := router.SetupRoutes()

	fmt.Println("Listening on port 8080 and printing Digital Ocean S3 URL " + os.Getenv("DIGITAL_OCEAN_S3_URL"))

	router.Run(":8080")
}

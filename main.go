package main

import (
	"memperbaikikode/routers"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routers.GetEngine()
	r.Run(":" + port)
}

package main

import (
	"binar-academy/example-db-rest-api/routers"
)

const SERVER_PORT string = ":8080"

func main() {
	r := routers.GetEngine()
	r.Run(SERVER_PORT)
}
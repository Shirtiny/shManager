package main

import (
	"shManager/server"
)

func main() {
	router := server.CreateRouter()
	router.Run(":2021")
}

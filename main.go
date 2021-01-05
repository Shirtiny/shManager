package main

import (
	"shManager/model"
	"shManager/server"
)

func main() {
	model.ConnectDatabase()
	router := server.CreateRouter()
	router.Run(":2021")
}

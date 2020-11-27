package main

import (
	"kmt1/config"
	"kmt1/handler"
)

func main() {
	config.LoadEnv()
	server := handler.NewServer()
	server.Start()
}

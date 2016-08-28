package main

import (
	"github.com/dragfire/gocha/client"
	"github.com/dragfire/gocha/socket"
)

func main() {
	socket.Connect()
	client.Init()
}

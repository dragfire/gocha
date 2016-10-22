package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dragfire/gocha/client"
	"github.com/dragfire/gocha/server"
)

var (
	errorNotEnoughArguments = errors.New("Supply a mode.")
	clientMode              = "client"
	serverMode              = "server"
)

func getMode() (string, error) {
	mode := os.Args
	n := len(mode)

	if n > 1 && (mode[1] == serverMode || mode[1] == clientMode) {
		return mode[1], nil
	}

	return "", errorNotEnoughArguments
}

func main() {

	// go run c.go client/backend

	mode, err := getMode()

	//fmt.Println("Mode:", mode, err)

	if err != nil {
		// create a panic in case mode is not provided
		panic(err)
	}

	fmt.Println("Mode:", mode)

	if mode == clientMode {
		client.Init()
	} else {
		server.Main()
	}

	// go func () {
	//     fmt.Println("Starting server")
	//     server.Main()
	// }()
	// defer client.Init()
}

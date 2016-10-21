package main 

import (
    "github.com/dragfire/gocha/client"
    "github.com/dragfire/gocha/server"
    "fmt"
    "os"
    "errors"
)

var (
    ErrorNotEnoughArguments = errors.New("Supply a mode.")
    CLIENT = "client"
    SERVER ="server"
)

func getMode() (string, error) {
    mode := os.Args
    n := len(mode)

    if n>1 && (mode[1] == SERVER || mode[1] == CLIENT)  {
        return mode[1], nil
    }

    return "", ErrorNotEnoughArguments
}

func main(){

    // go run c.go client/backend 

    mode, err := getMode()
    
    //fmt.Println("Mode:", mode, err)
    
    if err == nil {
        fmt.Println("Mode:", mode)
    } else {
        // create a panic in case mode is not provided
        panic(err)   
    }

    if mode == CLIENT {
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
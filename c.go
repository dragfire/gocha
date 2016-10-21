package main 

import (
    "github.com/dragfire/gocha/client"
    "github.com/dragfire/gocha/server"
    "fmt"
)

func main(){
    go func () {
        fmt.Println("Starting server")
        server.Main()
    }()
    defer client.Init()
}
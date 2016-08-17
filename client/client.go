package main

import (
	"bufio"
	"fmt"
	"github.com/dragfire/gocha/logger"
	"golang.org/x/net/websocket"
	"os"
)

var (
	url    = "ws://127.0.0.1:8080/socket"
	origin = "http://127.0.0.1:8080"
	in     = bufio.NewReader(os.Stdin)
)

func receiver(ws *websocket.Conn) {
	var msg string
	for {
		websocket.Message.Receive(ws, &msg)
		logger.Info.Println("Server sent: ", msg)
	}
}

func sender(ws *websocket.Conn) {
	var message string
	for {
		fmt.Print("Message: ")
		message, _ = in.ReadString('\n')
		websocket.Message.Send(ws, message)
	}
}

func connect() {
	ws, err := websocket.Dial(url, "", origin)
	ch := make(chan int)
	if err != nil {
		logger.Error.Println("Error", err)
	}

	//var msg string
	//websocket.Message.Receive(ws, &msg)
	//logger.Info.Println("Server sent: ", msg)

	go receiver(ws)
	go sender(ws)

	ch <- 1
	// if _, err := ws.Write([]byte("Hello World!\n")); err != nil {
	// 	logger.Error.Println("Error", err)
	// }
	//
	// var msg = make([]byte, 512)
	//
	// var n int
	//
	// if n, err = ws.Read(msg); err != nil {
	// 	logger.Error.Println(err)
	// }
	//
	// logger.Info.Println("Received: " + string(msg[:n]))
}

func main() {
	logger.Info.Println("******Starting Chat Client******")
	connect()
}

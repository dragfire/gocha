package socket

import (
	"golang.org/x/net/websocket"
	"log"
)

var (
	Conn   *websocket.Conn // setup a global socket connection for client
	url    = "ws://localhost:3000/socket"
	origin = "http://localhost/"
)

func receiver(ws *websocket.Conn) {
	var msg string
	for {
		websocket.Message.Receive(ws, &msg)
		log.Println("Server sent: ", msg)
	}
}

func sender(ws *websocket.Conn) {
	// var message string
	for {
		log.Println("sdegdg")
	}
}

func Connect() error {
	var err error
	log.Println("Websocket connection established.")
	Conn, err = websocket.Dial(url, "", origin)
	// ch := make(chan int)
	if err != nil {
		log.Panicln("Error:", err)
		return err
	}

	go receiver(Conn)
	go sender(Conn)
	return nil
}

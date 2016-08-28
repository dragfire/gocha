package socket

import (
	"golang.org/x/net/websocket"
	"log"
	_ "time"
)

type Socket struct {
	// socket id will be generated and sent from client side
	id        string
	room_name string
	conn      *websocket.Conn
}

type Room struct {
	name    string
	sockets map[string]Socket // {key: socket_id, value: socket_connection}
}

var Rooms = []Room{} // all rooms

func Handler(ws *websocket.Conn) {
	for {
		var reply string
		log.Println("Config: ", ws.Config())
		log.Println("LocalAddr: ", ws.LocalAddr().String())
		log.Println("RemoteAddr: ", ws.RemoteAddr().String())
		log.Println("Client: " + reply)
		msg := "Received: " + reply
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			log.Println("Socket error!", err)
			break
		}

		//logger.Info.Println("Sending: " + msg)
		if err := websocket.Message.Send(ws, msg); err != nil {
			log.Println("Can't send.", err)
			break
		}
	}
}

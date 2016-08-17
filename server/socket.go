package server

import (
	"github.com/dragfire/gocha/logger"
	"golang.org/x/net/websocket"
)

func SocketHandler(ws *websocket.Conn) {
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			logger.Error.Println("Socket error!", err)
			break
		}

		logger.Debug.Println("Config: ", ws.Config())
		logger.Debug.Println("LocalAddr: ", ws.LocalAddr().String())
		logger.Debug.Println("RemoteAddr: ", ws.RemoteAddr().String())
		logger.Info.Println("Client: " + reply)
		msg := "Received: " + reply
		//logger.Info.Println("Sending: " + msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			logger.Error.Println("Can't send.", err)
			break
		}
	}
}

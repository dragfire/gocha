package main

import (
	_ "fmt"
	"net/http"

	"github.com/dragfire/gocha/logger"
	middleware "github.com/dragfire/gocha/middlewares"
	"github.com/dragfire/gocha/server"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", middleware.LogRequest(http.HandlerFunc(server.Handler)))
	http.Handle("/view/", middleware.LogRequest(server.MakeHandler(server.ViewHandler)))
	http.Handle("/edit/", middleware.LogRequest(server.MakeHandler(server.EditHandler)))
	http.Handle("/save/", middleware.LogRequest(server.MakeHandler(server.SaveHandler)))
	http.Handle("/chat", middleware.LogRequest(http.HandlerFunc(server.ClientHandler)))
	http.Handle("/socket", websocket.Handler(server.SocketHandler))
	logger.Info.Println("Server Listening: 8080")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	logger.Debug.Println(http.ListenAndServe(":8080", nil))
}

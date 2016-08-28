package main

import (
	_ "fmt"
	"github.com/dragfire/gocha/logger"
	middleware "github.com/dragfire/gocha/middlewares"
	"github.com/dragfire/gocha/server"
	"github.com/dragfire/gocha/socket"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	http.Handle("/", middleware.LogRequest(http.HandlerFunc(server.Handler)))
	http.Handle("/view/", middleware.LogRequest(server.MakeHandler(server.ViewHandler)))
	http.Handle("/edit/", middleware.LogRequest(server.MakeHandler(server.EditHandler)))
	http.Handle("/save/", middleware.LogRequest(server.MakeHandler(server.SaveHandler)))
	http.Handle("/chat", middleware.LogRequest(http.HandlerFunc(server.ClientHandler)))
	http.Handle("/socket", websocket.Handler(socket.Handler))

	if PORT == "" {
		PORT = "3000"
	}
	logger.Info.Println("Server Listening:", PORT)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	logger.Debug.Println(http.ListenAndServe(":"+PORT, nil))
}

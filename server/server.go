package server

import (
	"github.com/dragfire/gocha/logger"
	"net/http"
	"os"
	"github.com/dragfire/gocha/middlewares"
	"golang.org/x/net/websocket"
	"github.com/dragfire/gocha/socket"
)

func Main() {
	PORT := os.Getenv("PORT")

	http.Handle("/", middleware.LogRequest(http.HandlerFunc(Handler)))
	http.Handle("/view/", middleware.LogRequest(MakeHandler(ViewHandler)))
	http.Handle("/edit/", middleware.LogRequest(MakeHandler(EditHandler)))
	http.Handle("/save/", middleware.LogRequest(MakeHandler(SaveHandler)))
	http.Handle("/chat", middleware.LogRequest(http.HandlerFunc(ClientHandler)))
	http.Handle("/socket", websocket.Handler(socket.Handler))

	if PORT == "" {
		PORT = "3000"
	}
	logger.Info.Println("Server Listening:", PORT)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	logger.Debug.Println(http.ListenAndServe(":"+PORT, nil))
	//client.Init()
}

package main

import (
	_ "fmt"
	"github.com/dragfire/gocha/logger"
	"github.com/dragfire/gocha/middlewares"
	"github.com/dragfire/gocha/server"
	"net/http"
)

func main() {
	//http.Handle("/", middleware.LogRequest(
	//http.HandlerFunc(
	//func(w http.ResponseWriter, r *http.Request) {
	//logger.Info.Println("Logging Done.")
	//})))
	//http.Handle("/", middleware.LogRequest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("OK!"))
	//})))
	http.Handle("/", middleware.LogRequest(http.HandlerFunc(server.Handler)))
	http.Handle("/view/", middleware.LogRequest(server.MakeHandler(server.ViewHandler)))
	http.Handle("/edit/", middleware.LogRequest(server.MakeHandler(server.EditHandler)))
	http.Handle("/save/", middleware.LogRequest(server.MakeHandler(server.SaveHandler)))
	logger.Info.Println("Server Listening: 8080")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	logger.Debug.Println(http.ListenAndServe(":8080", nil))
}

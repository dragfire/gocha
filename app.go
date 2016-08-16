package main

import (
	"fmt"
	"github.com/dragfire/gocha/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Handler)
	http.HandleFunc("/view/", server.MakeHandler(server.ViewHandler))
	http.HandleFunc("/edit/", server.MakeHandler(server.EditHandler))
	http.HandleFunc("/save/", server.MakeHandler(server.SaveHandler))
	fmt.Println("Server Listening: 8080")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

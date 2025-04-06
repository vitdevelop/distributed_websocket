package main

import (
	"distributed_websocket/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HandleDefaultPage)
	http.Handle("/{filename}", http.FileServer(http.Dir("./www")))
	http.Handle("/heroes/{filename}", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/ws", handler.HandleWs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

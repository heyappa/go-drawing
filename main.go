package main

import (
	"log"
	"net/http"
    "fmt"
//    "io/ioutil"
)

func main() {
	hub := newHub()
	go hub.run()

    http.Handle("/", http.FileServer(http.Dir("./client")))

	http.HandleFunc("/ws", hub.handleWebSocket)

    fmt.Println("starting web server (drawing) at 3000")
    err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

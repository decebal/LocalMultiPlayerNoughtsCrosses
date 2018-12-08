package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	host := ""
	port := os.Getenv("PORT")

	conn := fmt.Sprint(host, ":", port)
	router := http.NewServeMux()
	router.Handle("/ws", wsHandler{})

	log.Printf("Started Game on %v", conn)
	log.Fatal(http.ListenAndServe(conn, router))
}

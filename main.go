package main

import (
	"log"
	"net/http"

	web "web/functions"
	handler "web/handler"
)

func main() {
	web.FileReader("standard.txt")
	http.HandleFunc("/", handler.HandleAscii)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

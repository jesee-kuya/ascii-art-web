package main

import (
	"fmt"
	"log"
	"net/http"

	web "web/functions"
	handler "web/handler"
)

func main() {
	web.FileReader("standard.txt")
	http.HandleFunc("/", handler.HandleAscii)
	fmt.Println("App is running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

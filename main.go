package main

import (
	"fmt"
	"log"
	"net/http"

	handler "web/handler"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler.HandleAscii)
	fmt.Println("App is running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

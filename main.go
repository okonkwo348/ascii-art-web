package main

import (
	"ascii-art-web/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii-art", handler.AsciiHandler)
	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	port := ":8080"
	err := http.ListenAndServe(port, http.FileServer(http.Dir("apidoc")))
	if err != nil {
		log.Fatal("ListenAndServe fail:", err)
	}
}

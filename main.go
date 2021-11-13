package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":3200", router)
	log.Fatal(server)
}

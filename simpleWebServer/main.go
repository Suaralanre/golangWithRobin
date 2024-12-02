package main

import (
	"log"
	"net/http"
	"time"
)

var currentTime = time.Now()

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/get", GetTime)
	mux.HandleFunc("/set", SetTime)

	log.Println("Starting Server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}



package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var currentTime = time.Now()

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/get", GetTime)
	mux.HandleFunc("/set", SetTime)

	server := &http.Server{
		Addr: ":4000",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Println("Starting Server on port 4000")
		if err := server.ListenAndServe();
		err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP Server error: %v", err)
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down server....")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server exited properly")
}



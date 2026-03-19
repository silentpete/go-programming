package main

import (
	// "context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	srv := http.Server{
		Addr: "0.0.0.0:8080",
	}

	idleConnsClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	go func() {
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(nil); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()
	close(idleConnsClosed)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Printf("exiting")
}

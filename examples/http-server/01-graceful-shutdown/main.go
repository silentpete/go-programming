package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var address = ""
var port = "6060"

// rootHandler is answering on / ("root")
func rootHandler(w http.ResponseWriter, r *http.Request) {

	// display to browser request
	fmt.Fprintln(w, "hello world")

	// log who requests root
	log.Println(r.RemoteAddr, "requested", r.RequestURI)
}

func main() {

	// A Server defines parameters for running an HTTP server.
	// The zero value for Server is a valid configuration.
	log.Printf("creating new http server \"%v:%v\"\n", address, port)
	srv := http.Server{
		Addr: address + ":" + port,
	}

	log.Println("add rootHandler handler to http server.")
	http.HandleFunc("/", rootHandler)

	log.Println("make channel")
	ch := make(chan os.Signal, 1)
	log.Println("use channel to wait for signals")
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	log.Println("starting http server in goroutine")
	go func() {
		// ListenAndServe starts an HTTP server with a given address and handler.
		// The handler is usually nil, which means to use DefaultServeMux.
		// Handle and HandleFunc add handlers to DefaultServeMux:
		// ie. http.ListenAndServe(":6060", nil)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("ListenAndServe Error:", err)
		}
	}()

	log.Println("waiting for signal... (channels block)")
	s := <-ch
	log.Println("Got SIGNAL:", s)

	log.Println("close channel")
	close(ch)

	log.Println("shutdown http server")
	err := srv.Shutdown(context.Background())
	if err != nil {
		// Error from closing listeners, or context timeout
		log.Fatal("Shutdown error:", err)
	}

	log.Println("Supposed Graceful Shutdown")
}

// Reference
// https://golang.org/pkg/net/http/#Server
// https://golang.org/pkg/os/signal/#Notify
// https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2

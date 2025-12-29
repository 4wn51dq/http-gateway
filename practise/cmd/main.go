package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"practise/handlers"
	"time"
)

func main() {
	// Handle func is designed over the default serve multiplexer or and http request multiplexer
	// it uses the function to create an http handler and adds it to the defaultServeMUX
	//
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h := handlers.NewHandler(l)
	sh := handlers.NewSecondHandler(l)

	servemux := http.NewServeMux()
	servemux.Handle("/", h)
	servemux.Handle("/SecondHandler", sh)

	// the Server struct defines the parameters to run an http server
	// timeouts are protection parameters because request/response (write/read) takes time
	s := &http.Server{
		Addr:         ":8000",
		Handler:      servemux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	// we have created a goroutine for starting the go http serveer
	// for goroutines to pass around we need to build and use channels,
	// in the os, we get the signal package to use the channel to carry os signals

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

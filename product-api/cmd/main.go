package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	handlers "product-api/handlers"
	"time"

	"github.com/gorilla/mux"
	"k8s.io/utils/env"
)

var sec time.Duration = time.Second
var bindAddr = env.GetString("BIND_ADDRESS", ":8000")

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(l)
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter() // gives a route filtered specifically for http verb: GET
	getRouter.HandleFunc("/", productHandler.GetProducts)
	getRouter.Use(productHandler.MiddlewareProductValidation)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProducts)

	s := http.Server{
		Addr:         bindAddr,
		Handler:      serveMux,
		ErrorLog:     l,
		ReadTimeout:  5 * sec,
		WriteTimeout: 10 * sec,
		IdleTimeout:  15 * sec,
	}

	go func() {
		l.Println("listening and serving on localhost")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("couldnt start server, error: %s", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Printf("Get signal: %s", sig)

	ctx, _ := context.WithTimeout(context.Background(), 10*sec)
	err := s.Shutdown(ctx)
	if err != nil {
		l.Println("couldnt execute graceful shutdown")
	}
}

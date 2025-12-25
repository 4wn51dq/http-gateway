package main

import (
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r chi.Mux = chi.NewRouter() // chi.Mux is only a struct to set up api
	handlers.Handler(r)
	fmt.Println("Starting go api service..")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}

package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// an http handler is an interface with a method to it: ServeHTTP!

type Handler struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
	// its important to log things sometimes because while
	// we are connecting to database and other services, we need
	// to write positive unit tests.
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.l.Println("enrguonewuio")
	// body is a i/o read closer
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "noooo", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "value from request header is: %d", d)

}

type SecondHandler struct {
	l *log.Logger
}

func NewSecondHandler(l *log.Logger) *SecondHandler {
	return &SecondHandler{l}
}

func (sh *SecondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

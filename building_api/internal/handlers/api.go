package handlers

import (
	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		// middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}

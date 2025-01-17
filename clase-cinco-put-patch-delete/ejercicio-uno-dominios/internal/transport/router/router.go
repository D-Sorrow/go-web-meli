package router

import "github.com/go-chi/chi/v5"

func NewRouter(subs map[string]*chi.Mux) *chi.Mux {

	router := chi.NewRouter()
	for dir, sub := range subs {
		router.Mount(dir, sub)
	}
	return router
}

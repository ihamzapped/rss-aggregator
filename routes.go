package main

import "github.com/go-chi/chi/v5"

func (api *ApiConfig) initRoutes() *chi.Mux {

	router := chi.NewRouter()
	v1router := chi.NewRouter()

	v1router.Post("/users", api.handleCreateUser)
	v1router.Get("/users", api.middlewareAuth(api.handleGetUser))

	v1router.Post("/feeds", api.middlewareAuth(api.handleCreateFeed))

	v1router.Get("/healthz", healthCheck)

	v1router.Get("/err", errCheck)

	router.Mount("/v1", v1router)

	return router

}

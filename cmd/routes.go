package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *ApiConfig) initRoutes() *chi.Mux {

	router := chi.NewRouter()
	v1router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	v1router.Post("/user", api.handleCreateUser)
	v1router.Get("/user", api.middlewareAuth(api.handleGetUser))

	// For demo purposes
	v1router.Get("/users", api.handleGetAllUsers)

	v1router.Get("/feeds", api.handleGetAllFeeds)
	v1router.Post("/feeds", api.middlewareAuth(api.handleCreateFeed))

	v1router.Post("/feed-follows", api.middlewareAuth(api.handleCreateFeedFollow))
	v1router.Get("/feed-follows", api.middlewareAuth(api.handleGetUserFollows))
	v1router.Delete("/feed-follows/{follow_id}", api.middlewareAuth(api.handleDeleteUserFollow))

	v1router.Get("/user-feeds", api.middlewareAuth(api.handleGetPostsForUser))

	v1router.Get("/healthz", healthCheck)

	v1router.Get("/err", errCheck)

	router.Mount("/v1", v1router)

	return router

}

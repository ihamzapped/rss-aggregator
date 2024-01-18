package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("No port defined in the env")
	}

	router := chi.NewRouter()
	v1router := chi.NewRouter()

	v1router.Get("/healthz", healthCheck)

	v1router.Get("/err", errCheck)

	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server started on port: %v", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

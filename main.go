package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/ihamzapped/rss-aggregator/internal/database"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not defined in the env")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL not defined in the env")

	}

	dbConn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	api := &ApiConfig{
		DB: database.New(dbConn),
	}

	router := chi.NewRouter()
	v1router := chi.NewRouter()

	v1router.Post("/users", api.handleCreateUser)
	v1router.Get("/users", api.handleGetUser)

	v1router.Get("/healthz", healthCheck)

	v1router.Get("/err", errCheck)

	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Print("Server started on port: ", port)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

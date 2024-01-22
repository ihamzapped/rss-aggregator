package main

import (
	"database/sql"
	"github.com/ihamzapped/rss-aggregator/internal/database"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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

	server := &http.Server{
		Handler: api.initRoutes(),
		Addr:    ":" + port,
	}

	log.Print("Server started on port: ", port)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

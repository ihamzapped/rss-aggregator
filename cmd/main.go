package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/ihamzapped/rss-aggregator/internal/database"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	router := api.initRoutes()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Print("Server started on port: ", port)

	go startScraping(api.DB, 10, time.Minute)

	// Create a route along / that will serve contents from
	// the ./public/ folder.
	workDir, _ := os.Getwd()
	staticDir := http.Dir(filepath.Join(workDir, "public"))

	fileServer(router, "/", staticDir)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

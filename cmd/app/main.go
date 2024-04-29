package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dimitar-ivanov-93/htmx/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := chi.NewMux()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Request URL:", r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})
	routes.SetupRoutes(router)

	listenAddr := os.Getenv("LISTEN_ADDR")
	log.Printf("Starting server on %s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}

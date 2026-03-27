package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/shivsperfect/uwe/handler"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := chi.NewMux()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/customer/{id}", handler.ServeHTTP(handler.HandleGetCustomer))

	r.Post("/upload", handler.ServeHTTP(handler.HandleUpload))

	port := os.Getenv("PORT")
	slog.Info("Listening on PORT %s", port)
	http.ListenAndServe(port, r)
}

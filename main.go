package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/shivsperfect/uwe/db"
	"github.com/shivsperfect/uwe/handler"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.Create()
	uploadHandler := handler.NewUploadHandler(db)

	router := chi.NewMux()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/customer/{id}", handler.ServeHTTP(handler.HandleGetCustomer))

	router.Post("/file", handler.ServeHTTP(uploadHandler.HandleCreateFileUpload))
	router.Post("/file/{id}", handler.ServeHTTP(uploadHandler.HandleFileUpload))

	port := os.Getenv("PORT")
	slog.Info("Listening on PORT %s", port)
	http.ListenAndServe(port, router)
}

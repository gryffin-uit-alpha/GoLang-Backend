package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	fmt.Println("Hello World")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	} else {
		fmt.Println("Port is: ", portString)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		data := Response{
			Message: "Your API is Healthy and Working Correctly",
			Status:  200,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

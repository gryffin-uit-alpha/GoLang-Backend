package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gryffin-uit-alpha/GoLang-Backend/internal/db"
	handlers "github.com/gryffin-uit-alpha/GoLang-Backend/internal/handler"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	apiCfg := &apiConfig{
		DB: db.New(conn),
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

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlers.HandlerHealth)

	v1Router.Post("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(apiCfg.DB, w, r)
	})
	v1Router.Post("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginUser(apiCfg.DB, w, r)
	})

	v1Router.Post("/categories", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateCategory(apiCfg.DB, w, r)
	})
	v1Router.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListCategories(apiCfg.DB, w, r)
	})

	v1Router.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePost(apiCfg.DB, w, r)
	})
	v1Router.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListPosts(apiCfg.DB, w, r)
	})
	v1Router.Get("/posts/{slug}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPost(apiCfg.DB, w, r)
	})
	v1Router.Delete("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePost(apiCfg.DB, w, r)
	})

	router.Mount("/api/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

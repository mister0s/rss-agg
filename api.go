package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mister0s/rss-agg/internal/database"
)

type ApiServer struct {
	port string
	DB   *database.Queries
}

func NewApiServer(port string, DB *database.Queries) ApiServer {
	return ApiServer{
		port: port,
		DB:   DB,
	}
}

func (s ApiServer) Run() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http", "https"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health-check", healthCheckHandler)
	v1Router.Get("/error", errHandler)
	v1Router.Get("/user", s.getUserHandler)
	v1Router.Post("/user", s.createUserHandler)

	router.Mount("/v1", v1Router)
	server := &http.Server{
		Addr:    ":" + s.port,
		Handler: router,
	}
	log.Printf("Server running on %v", s.port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("PORT: ", err)
	}
}

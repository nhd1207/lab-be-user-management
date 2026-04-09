package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/nhd1207/be-user-management/internal/repository/postgres"
	"github.com/nhd1207/be-user-management/internal/service"
	handler "github.com/nhd1207/be-user-management/internal/transport/http"
)

func main() {
	log.Println("Starting server...")
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db := postgres.NewDB(dbURL)

	repo := postgres.NewUserRepository(db)

	hasher := &service.BcryptHasher{}
	userService := service.NewUserService(repo, hasher)
	handler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", handler.CreateUser)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

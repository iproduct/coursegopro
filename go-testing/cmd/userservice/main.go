package main

import (
	"github.com/EmilGeorgiev/training/go-testing/userservice/api"
	"github.com/EmilGeorgiev/training/go-testing/userservice/persistent"
	"github.com/clouway/godb/datastoretest"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := datastoretest.NewDatabase()
	defer db.Close()

	repo := persistent.NewUserRepository(db)
	userHandler := api.UserHandler{UserRepo: repo}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/users/{id}", userHandler.Get)
	r.Post("/users", userHandler.Create)

	_ = http.ListenAndServe(":3000", r)
}

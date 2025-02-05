package users

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func UserHandlers() http.Handler {
	r := chi.NewRouter()

	r.Get("/v1/api/users", GetUsersHandler)
	r.Get("/v1/api/users/{userId}", GetUserByIdHandler)
	r.Post("/v1/api/users", CreateUserHandler)

	return r
}

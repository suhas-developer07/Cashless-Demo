package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes()http.Handler {
	r := chi.NewRouter()

	return r

}
package router

import (
	"net/http"

	"github.com/MuriLovin/fin-control/internal/handlers"
	"github.com/MuriLovin/fin-control/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetUsersRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/create", middlewares.LoggingMiddleware(handlers.SaveUser)).Methods(http.MethodPost)
	router.HandleFunc("/users/{userId}", middlewares.LoggingMiddleware(handlers.FindUser)).Methods(http.MethodGet)
	return router
}

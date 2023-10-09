package router

import (
	"net/http"

	"github.com/MuriLovin/fin-control/internal/handlers"
	"github.com/MuriLovin/fin-control/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetManagementsRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/management/all/user/{userId}", middlewares.LoggingMiddleware(handlers.FindAllUserManagement)).Methods(http.MethodGet)
	router.HandleFunc("/management/create", middlewares.LoggingMiddleware(handlers.SaveManagement)).Methods(http.MethodPost)
	router.HandleFunc("/management/{managementId}", middlewares.LoggingMiddleware(handlers.FindManagement)).Methods(http.MethodGet)
	return router
}

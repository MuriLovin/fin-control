package router

import (
	"net/http"

	"github.com/MuriLovin/fin-control/internal/handlers"
	"github.com/MuriLovin/fin-control/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetCategoriesRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/category/create", middlewares.LoggingMiddleware(handlers.SaveCategory)).Methods(http.MethodPost)
	router.HandleFunc("/category/{categoryId}", middlewares.LoggingMiddleware(handlers.FindCategory)).Methods(http.MethodGet)
	return router
}

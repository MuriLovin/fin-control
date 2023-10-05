package router

import (
	"github.com/MuriLovin/fin-control/internal/handlers"
	"github.com/MuriLovin/fin-control/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetUsersRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", middlewares.LoggingMiddleware(handlers.AllUsers)).Methods("GET")
	router.HandleFunc("/users", middlewares.LoggingMiddleware(handlers.SaveUser)).Methods("POST")
	router.HandleFunc("/hello/{friend}", middlewares.LoggingMiddleware(handlers.FriendHandler))
	return router
}

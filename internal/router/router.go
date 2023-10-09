package router

import "github.com/gorilla/mux"

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	SetUsersRouter(router)
	SetCategoriesRouter(router)
	SetManagementsRouter(router)
	return router
}

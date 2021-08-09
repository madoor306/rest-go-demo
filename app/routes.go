package main

import "github.com/gorilla/mux"

// Initialiaze all the dispatch handlers and return the max router
func (handler *DemoApp) routes() *mux.Router {

	router := mux.NewRouter()

	// application routes
	router.HandleFunc("/v1/users", handler.getAllUsers).Methods("GET")
	router.HandleFunc("/v1/users/{user_id}", handler.getUser).Methods("GET")
	router.HandleFunc("/v1/users", handler.createUser).Methods("POST")
	router.HandleFunc("/v1/users/{user_id}", handler.updateUser).Methods("PUT")
	router.HandleFunc("/v1/users/{user_id}", handler.deleteUser).Methods("DELETE")

	return router
}

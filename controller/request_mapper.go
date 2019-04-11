package controller

import "github.com/gorilla/mux"

// MapHandlers function does the mapping from the request URLs to the correspondent handler functions
func MapHandlers() (*mux.Router, error) {
	router := mux.NewRouter()
	context := router.PathPrefix("/v1").Subrouter()
	context.Methods("OPTIONS").HandlerFunc(corsHandler)
	context.Methods("GET").Path("/health").HandlerFunc(checkHealth)
	return router, nil
}

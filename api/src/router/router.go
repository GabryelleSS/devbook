package router

import "github.com/gorilla/mux"

// returns the configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
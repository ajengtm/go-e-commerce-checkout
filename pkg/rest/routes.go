package rest

import (
	"github.com/gorilla/mux"
)

// SetupRoutes initializes the REST API routes.
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cart/{userID}/add", AddItemToCart).Methods("POST")
	r.HandleFunc("/cart/{userID}/checkout", Checkout).Methods("GET")
	r.HandleFunc("/cart/{userID}/pay", Pay).Methods("POST")
	return r
}

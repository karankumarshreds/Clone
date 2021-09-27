package rest 

import (
	"github.com/gorilla/mux"
)

func InitHandler() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/", welcomeHandler).Methods("GET")
	return router
}
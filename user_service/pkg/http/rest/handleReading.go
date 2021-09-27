package rest 

import (
	"net/http"
	"encoding/json"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the app!")
}
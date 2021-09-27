package main

import (
	"net/http"
	"log"
	"github.com/karankumarshreds/Clone/user_service/pkg/http/rest"
)

func main() {
	log.Println("Starting on port 8000")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8000", router))

}


// package main 

// import (
// 	"github.com/karankumarshreds/Clone/user_service/cmd/app"
// ) 

// func main() {

// 	a := app.App{}
// 	a.Initialize(
// 		"8000",
// 		"development",
// 		app.DB_CONFIG{
// 			User: "user_service",
// 			Database: "users_db",
// 			Password: "password",
// 		},
// 	)
// 	a.Run()

// }


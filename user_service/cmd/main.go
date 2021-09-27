package main

import (
	"log"
	"http/rest"
)

func main() {
	log.Println("Starting on port 8000")
	router := rest.initHandlers

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


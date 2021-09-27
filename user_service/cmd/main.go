package main 

import (
	"github.com/karankumarshreds/Clone/user_service/cmd/app"
) 

func main() {
	a := app.App{}
	a.Initialize("8000", "development", app.DB_CONFIG{user: "user_service", password: "password", database: "users_db"})
	a.Run()
}


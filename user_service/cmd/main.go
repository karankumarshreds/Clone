package main 

func main() {
	 
	app := App{}
	// .env 
	// 
	//
	//
	
	app.Initialize("8001", "development", DB_CONFIG{database: "users_db", user: "user_service", password:"password"})
	app.Run()

}
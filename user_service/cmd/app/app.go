package app 

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
)

type App struct {

	infoLog *log.Logger
	errorLog *log.Logger 
	
	DB *sql.DB	

	config struct {
		port string 
		env  string 
	}

}

type DB_CONFIG struct {
	User string
	Password string 
	Database string 
}


func (app *App) Initialize(PORT string, ENV string, DB_CONFIG DB_CONFIG) {

	// logger 
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	app.infoLog = infoLog
	app.errorLog = errorLog

	// app 
	app.config.env = ENV 
	app.config.port = PORT 

	// database 
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_CONFIG.User, DB_CONFIG.Password, DB_CONFIG.Database)
	DB, err := sql.Open("postgres", connectionString)	
	app.failOnError(err, "Unable to connect to the database")
	app.infoLog.Println("Connection made with Database")
	app.DB = DB
		
}

func (app *App) Run() {
	router := mux.NewRouter()
	app.infoLog.Printf("Starting app in %s mode on port %s", app.config.env, app.config.port)
	err := http.ListenAndServe(fmt.Sprintf(":" + app.config.port), router)
	app.failOnError(err, "Unable to run the application")
}

func (app *App) failOnError(err error, msg string) {
	if err != nil {
		app.errorLog.Fatalf(msg + "%s\t\n", err)
	}
}
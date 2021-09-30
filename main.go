package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/render", render).Methods("POST")
	app.Router.HandleFunc("/render", getRender).Methods("GET")
	app.Router.HandleFunc("/", helloWorldHandler).Methods("GET")
	app.Router.HandleFunc("/health", getHealth).Methods("GET")
}

func (app *App) run() {

	config, err := newConfig("./config.yml")
	ExitIfError(err)
	if config.Server.Verbosity != "info" {
		log.Println("Host           : " + config.Server.Host)
		log.Println("Port           : " + config.Server.Port)
	}
	log.Println("#########")
	log.Fatal(http.ListenAndServe(config.Server.Host+":"+config.Server.Port, app.Router))
}

func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}

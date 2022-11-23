package app

import (
	"net/http"
	"sensors/app/controllers"
	"sensors/app/services"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.routes()
	go services.ListenSensors()

	return a
}

func (a *App) routes() {
	a.Router.HandleFunc("/", controllers.Index()).Methods("GET")
	a.Router.HandleFunc("/sensor", controllers.GetSensors()).Methods("GET")
	a.Router.HandleFunc("/sensor", controllers.AddSensor()).Methods("POST")
	a.Router.HandleFunc("/sensor/{id}", controllers.DeleteSensor()).Methods("DELETE")
	a.Router.HandleFunc("/sensor/{id}", controllers.UpdateSensor()).Methods("PUT")

	a.Router.NotFoundHandler = a.Router.NewRoute().HandlerFunc(http.NotFound).GetHandler()
}

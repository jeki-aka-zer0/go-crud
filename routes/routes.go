package routes

import (
	"github.com/jeki-aka-zer0/go-crud/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage).Methods("GET")
	router.HandleFunc("/engineer", controllers.ReadAllEngineers).Methods("GET")
	router.HandleFunc("/engineer/{id}", controllers.GetEngineerById).Methods("GET")
	router.HandleFunc("/engineer", controllers.CreateEngineer).Methods("PUT")
}

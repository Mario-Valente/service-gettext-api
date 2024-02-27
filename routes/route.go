package routes

import (
	"service-gettext-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigRoutes(r *mux.Router) {
	r.HandleFunc("/total-string/", controllers.CountCaracter).Methods("POST")
}

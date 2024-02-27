package main

import (
	"fmt"
	"net/http"
	"service-gettext-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.ConfigRoutes(r)
	fmt.Println("Servidor iniciado na porta 8080...")
	http.ListenAndServe(":8080", r) // Passando o roteador r como o segundo argumento
}

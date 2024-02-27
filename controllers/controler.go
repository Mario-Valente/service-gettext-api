package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service-gettext-api/models"
)

// func TotalStringHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	ParameterToString := params["total-Parameter"]
// }

func CountCaracter(w http.ResponseWriter, r *http.Request) {
	var text models.Text
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numCaracteres := len(text.Content)
	fmt.Fprintf(w, "O texto tem %d caracteres.\n", numCaracteres)
}

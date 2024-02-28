package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service-gettext-api/models"
	"strings"
)

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

func ToUpper(w http.ResponseWriter, r *http.Request) {
	var text models.Text
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	upperCase := strings.ToUpper(text.Content)
	fmt.Fprint(w, "text upper: \n", upperCase)
}

func ToLower(w http.ResponseWriter, r *http.Request) {
	var text models.Text
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lowerCase := strings.ToLower(text.Content)
	fmt.Fprint(w, "text lower: \n", lowerCase)
}

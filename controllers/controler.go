package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"service-gettext-api/models"
	"strconv"
	"strings"
	"time"
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

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		"!@#$%^&*()_+"

	rand.Seed(time.Now().UnixNano())
	var randomString []byte
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		randomString = append(randomString, charset[randomIndex])
	}
	return string(randomString)
}

func RandomStringHandler(w http.ResponseWriter, r *http.Request) {
	lengthStr := r.URL.Query().Get("length")
	if lengthStr == "" {
		lengthStr = "12"
	}

	length, err := strconv.Atoi(lengthStr)
	if err != nil || length <= 0 {
		http.Error(w, "Invalid length parameter", http.StatusBadRequest)
		return
	}
	randomString := GenerateRandomString(length)
	response := map[string]string{"randomString": randomString}
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

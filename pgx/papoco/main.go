package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}



func HandleInputTelegram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%#v\n", data)
}

func main() {

	http.HandleFunc("/", HandleInputTelegram)

	log.Fatalln(http.ListenAndServe(":3000", nil))
}

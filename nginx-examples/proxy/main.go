package main

import (
	"log"
	"net/http"
	"os"
)

func App(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	servant, _ := os.Hostname()

	w.Write([]byte("Welcome to api home !!!" + servant))
}

func main() {

	http.HandleFunc("/app", App)

	log.Println("server running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}

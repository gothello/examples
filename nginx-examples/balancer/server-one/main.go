package main

import (
	"log"
	"net/http"
	"os"
)

func serverOne(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server One " + h))
}

func main() {

	http.HandleFunc("/", serverOne)

	log.Println("Server running port 3000")

	log.Fatalln(http.ListenAndServe(":3000", nil))
}

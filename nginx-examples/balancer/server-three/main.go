package main

import (
	"log"
	"net/http"
)

func serverThree(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Three"))
}

func main() {

	http.HandleFunc("/", serverThree)

	log.Println("Server running port 3000")

	log.Fatalln(http.ListenAndServe(":3000", nil))
}

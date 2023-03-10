package main

import (
	"log"
	"net/http"
)

func serverTwo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Two"))
}

func main() {

	http.HandleFunc("/", serverTwo)

	log.Println("Server running port 3000")

	log.Fatalln(http.ListenAndServe(":3000", nil))
}

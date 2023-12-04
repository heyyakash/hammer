package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	log.Print("Sshh!! Server is listening")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal("Server cannot start!")
	}

}

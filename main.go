package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}
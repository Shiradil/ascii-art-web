package cmd

import (
	"ascii-art-web/server/handlers"
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	http.Handle("/client/templates/", http.StripPrefix("/client/templates/", http.FileServer(http.Dir("./client/templates"))))
	port := ":8080"
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/generate", handlers.GenerateHandler)

	fmt.Printf("Starting server on localhost%s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

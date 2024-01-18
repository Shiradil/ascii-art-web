package cmd

import (
	"ascii-art-web/server/handlers"
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	http.Handle("/client/templates/", http.StripPrefix("/client/templates/", http.FileServer(http.Dir("./client/templates"))))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/generate", handlers.GenerateHandler)
	
	port := ":8080"
	fmt.Printf("Starting server on localhost%s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

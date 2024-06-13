package main

import (
	"log"
	"net/http"

    "main/internal/handlers"
)

func main() {
    http.HandleFunc("/", handlers.MainPageHandler)
    http.HandleFunc("/guess", handlers.GuessHandler)
    http.HandleFunc("/realWord", handlers.RealWordHandler)

    fs := http.FileServer(http.Dir("src"))
    http.Handle("/src/", http.StripPrefix("/src/", fs))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

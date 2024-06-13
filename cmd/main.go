package main

import (
	"log"
	"net/http"

    "main/internal/handlers"
    "main/internal/logic"
)

func main() {
    port := ":8080"
    logic.SetRandomWord()

    http.HandleFunc("/", handlers.MainPageHandler)
    http.HandleFunc("/guess", handlers.GuessHandler)
    http.HandleFunc("/realWord", handlers.RealWordHandler)

    fs := http.FileServer(http.Dir("../src"))
    http.Handle("/src/", http.StripPrefix("/src/", fs))

    log.Println("Listening on port (", port, ")")
    log.Fatal(http.ListenAndServe(port, nil))
}

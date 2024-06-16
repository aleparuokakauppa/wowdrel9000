package main

import (
	"log"
	"net/http"
    "fmt"

    "main/internal/handlers"
    "main/internal/logic"
)

func main() {
    // Add check that looks at CWD
    // It has to be the root of the project
    config, err := handlers.GetConfig()
    if err != nil {
        log.Fatal(err)
    }

    port := fmt.Sprintf(":%d", config.Server.Port)

    http.HandleFunc("/", handlers.MainPageHandler)
    http.HandleFunc("/guess", handlers.GuessHandler)
    http.HandleFunc("/realWord", handlers.RealWordHandler)

    fs := http.FileServer(http.Dir("../src"))
    http.Handle("/src/", http.StripPrefix("/src/", fs))

    log.Println("Listening on port (", port, ")")
    log.Fatal(http.ListenAndServe(port, nil))
}

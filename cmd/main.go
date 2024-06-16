package main

import (
	"log"
	"net/http"
    "fmt"
    "os"

    "main/internal/httpHandlers"
    "main/internal/config"
    "main/internal/logic"
)

func main() {
    _, err := os.Stat("cmd/main.go")
    if err != nil {
        log.Fatal("The working dir should be run at the root of the project!")
    }

    config, err := config.GetConfig()
    if err != nil {
        log.Fatalf("Failed to load config file: %v", err)
    }

    go logic.SetRandomWordLoop(config.Server.WordRotateInterval)

    port := fmt.Sprintf(":%d", config.Server.Port)

    http.HandleFunc("/", handlers.MainPageHandler)
    http.HandleFunc("/guess", handlers.GuessHandler)
    http.HandleFunc("/realWord", handlers.RealWordHandler)

    fs := http.FileServer(http.Dir("./src"))
    http.Handle("/src/", http.StripPrefix("/src/", fs))

    log.Println("Listening on port (", port, ")")
    log.Fatal(http.ListenAndServe(port, nil))
}

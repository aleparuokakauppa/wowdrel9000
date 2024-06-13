package handlers

import (
    "net/http"
    "io"
    "encoding/json"
    "log"

    "main/internal/types"
    "main/internal/logic"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../src/index.html")
}

func GuessHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusBadRequest)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Unable to read request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    var guess types.Guess
    err = json.Unmarshal(body, &guess) 
    if err != nil {
        http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
        return
    }

    if guess.Version != 1 {
        http.Error(w, "Wrong protocol version. Want version 1.", http.StatusBadRequest)
    }

    letters := logic.CompareGuess(guess)

    response := types.GuessResponse{
        Version: 1,
        Win: true,
        Letters: letters,
    }

    // Check for negative win
    for _, letter := range letters {
        if letter.Status == "miss" || letter.Status == "close" {
            response.Win = false
        }
    }

    responseData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Server failed to marshal response JSON", http.StatusInternalServerError)
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseData)
}

func RealWordHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusBadRequest)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Unable to read request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    var guess types.Guess
    err = json.Unmarshal(body, &guess) 
    if err != nil {
        http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
        return
    }

    if guess.Version != 1 {
        http.Error(w, "Wrong protocol version. Want version 1.", http.StatusBadRequest)
    }

    result, err := logic.CheckRealWord(guess.Guess)
    if err != nil {
        log.Fatal(err)
        http.Error(w, "Server couldn't parse the file for real words.", http.StatusInternalServerError)
    }
    
    response := types.RealWordResponse{
        Version: 1,
        IsReal: result,
    }

    responseData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Server failed to marshal response JSON", http.StatusInternalServerError)
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseData)
}

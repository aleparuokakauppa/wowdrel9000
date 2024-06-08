package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var answer string

type Guess struct {
    Version int `json:"version"`
    Guess string `json:"guess"`
}

type Letter struct {
    Char rune `json:"char"`
    Status string `json:"status"`
}

type GuessResponse struct {
    Version int `json:"version"`
    Win bool `json:"win"`
    Letters [5]Letter `json:"letters"`
}

func servePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "web/index.html")
}

func getWords(infile string) ([]string, error) {
    wordFile, err := os.Open(infile)
    if err != nil {
        return nil, err
    }
    defer wordFile.Close()

    var words []string

    scanner := bufio.NewScanner(wordFile)
    for scanner.Scan() {
        line := scanner.Text()
        if (len(line) != 5) {
            return nil, fmt.Errorf("File contained word that wasn't 5 chars long!: %s", line)
        }
        words = append(words, strings.ToUpper(line))
    }

    return words, nil
}

func gameServerHandler() {
    words, err := getWords("words.txt")
    if err != nil {
        log.Fatal(err)
    }
    answer = words[rand.Intn(len(words))]
}

func compareGuess(guess Guess) [5]Letter {
    letters := new([5]Letter)

    for guessIndex, guessRune := range guess.Guess {
        letters[guessIndex].Char = guessRune
        guessMatch := false
        guessClose := false
        for answerIndex, answerRune := range answer {
            if guessRune == answerRune {
                if guessIndex == answerIndex {
                    guessMatch = true;
                }
                guessClose = true;
            }
        }
        if guessMatch == true {
            letters[guessIndex].Status = "match"
        } else if guessClose == true {
            letters[guessIndex].Status = "close"
        } else {
            letters[guessIndex].Status = "miss"
        }
    }
    return *letters
}

func handleGuess(w http.ResponseWriter, r *http.Request) {
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

    var guess Guess
    err = json.Unmarshal(body, &guess) 
    if err != nil {
        http.Error(w, "Unable to parse JSON data", http.StatusBadRequest)
        return
    }

    if guess.Version != 1 {
        http.Error(w, "Wrong protocol version. Want version 1.", http.StatusBadRequest)
    }

    letters := compareGuess(guess)

    response := GuessResponse{
        Version: 1,
        Win: true,
        Letters: letters,
    }

    // Check for negative win
    for _, letter := range letters {
        if letter.Status == "miss" {
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

func main() {
    http.HandleFunc("/", servePage)
    http.HandleFunc("/guess", handleGuess)

    fs := http.FileServer(http.Dir("web"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    go gameServerHandler()
    log.Fatal(http.ListenAndServe(":8080", nil))
}


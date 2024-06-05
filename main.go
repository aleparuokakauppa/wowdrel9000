package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

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

    fileBuf := bufio.NewReader(wordFile)
    for line, err := fileBuf.ReadString('\n'); err != io.EOF; {
        if (len(line) != 5) {
            return nil, fmt.Errorf("File contained word that wasn't 5 chars long!: %s", line)
        }
        words = append(words, line)
    }

    return words, nil
}

func gameServerHandler() {
    words, err := getWords("words.txt")
    if err != nil {
        log.Fatal(err)
    }

    randWord := words[rand.Intn(len(words))]
}

func main() {
    http.HandleFunc("/", servePage)

    fs := http.FileServer(http.Dir("web"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    go gameServerHandler()
    log.Fatal(http.ListenAndServe(":8080", nil))
}


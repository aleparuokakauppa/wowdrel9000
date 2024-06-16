package logic

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"main/internal/db"
	"main/internal/types"
    "main/internal/handlers"

    /* TODO break the import cycle with refactoring */

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var answer string
const WordFile string =  "../data/words.txt"

func CheckRealWord(clientWord string) (bool, error){
    wordFile, err := os.Open(WordFile)
    if err != nil {
        return false, fmt.Errorf("Cannot open wordfile: %v", err)
    }
    defer wordFile.Close()
    scanner := bufio.NewScanner(wordFile)
    for scanner.Scan() {
        line := scanner.Text()
        if (clientWord == strings.ToUpper(line)) {
            return true, nil
        }
    }
    return false, nil
}

// Runs the SetRandomWord method according to
// set interval.
func SetRandomWordLoop(interval int) {
    for {
        SetRandomWord()
        nextTime := time.Now().Add(time.Duration(interval) * time.Hour)
        duration := nextTime.Sub(time.Now())
        time.Sleep(duration)
    }
}

func SetRandomWord() {
    words, err := GetWords()
    if err != nil {
        log.Fatal(err)
    }
    answer = words[rand.Intn(len(words))]
    log.Println("Answer set to ", answer)
}

func GetWords() ([]string, error) {
    configClient, err := handlers.GetConfig()
    if err != nil {
        return []string{}, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configClient.Server.WordRotateInterval) * time.Second)
    defer cancel()

    dbClient, err := db.GetMongoClient(ctx)
    if err != nil {
        return []string{}, err
    }


    wordFile, err := os.Open(WordFile)
    if err != nil {
        return nil, fmt.Errorf("GetWords | Cannot open wordfile: %v", err)
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

func CompareGuess(guess types.Guess) [5]types.Letter {
    letters := new([5]types.Letter)

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

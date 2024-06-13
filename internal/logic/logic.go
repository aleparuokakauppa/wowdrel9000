package logic

import (
    "log"
    "os"
    "bufio"
    "strings"
    "math/rand"
    "fmt"

    "main/internal/types"
)

var answer string

func CheckRealWord(clientWord string ,infile string) (bool, error){
    wordFile, err := os.Open(infile)
    if err != nil {
        log.Println("Can't open wordfile: ", err)
        return false, err
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

func SetRandomWord(infile string) {
    words, err := GetWords(infile)
    if err != nil {
        log.Fatal(err)
    }
    answer = words[rand.Intn(len(words))]
    log.Println("Answer set to ", answer)
}

func GetWords(infile string) ([]string, error) {
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

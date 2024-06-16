import GuessLetterStack from '/src/js/guessLetterStack.js';
import { handleKeyEvent } from '/src/js/keypressUtils.js';
import { drawTileChars, drawTileColors } from '/src/js/tileUtils.js';
import { checkRealWord, checkWin } from '/src/js/apiClients.js'
import { smallNotify, notifyPlayerWin, notifyPlayerLoss } from '/src/js/notify.js';
import { drawLetterColors } from '/src/js/keypadUtils.js';


var currRow = 1;
var maxRows = 6;
var win = false;
var guessedWords = []

var guessLetterStack = new GuessLetterStack(5)

let getCurrRowObj = () => {
    return document.getElementById('row-' + currRow);
}

async function handleGuess() {
    let alreadyGuessed = (guessedWords.indexOf(guessLetterStack.toString()) > -1)
    if (alreadyGuessed === true) {
        guessLetterStack = new GuessLetterStack(5)
        drawTileChars(guessLetterStack, getCurrRowObj())
        smallNotify("Word is already guessed.")
        return false;
    }

    let isRealWord = await checkRealWord(guessLetterStack.toString());
    if (isRealWord === false) {
        guessLetterStack = new GuessLetterStack(5)
        drawTileChars(guessLetterStack, getCurrRowObj())
        smallNotify("Not a real word.")
        return false;
    }

    // This includes guess accuracy per letter
    let winResponseJSON = await checkWin(guessLetterStack.toString());

    drawTileColors(winResponseJSON, getCurrRowObj())
    drawLetterColors(winResponseJSON)
    guessLetterStack.clear();

    win = winResponseJSON.win;
    return true;
}

async function gameHandler() {
    // Main game loop
    while (currRow <= maxRows && win === false) {

        drawTileChars(guessLetterStack, getCurrRowObj());

        guessLetterStack = await handleKeyEvent(guessLetterStack);

        if (guessLetterStack.guess) {
            let word = guessLetterStack.toString();
            let nextGuess = await handleGuess();
            if (nextGuess === true) {
                guessedWords.push(word)
                currRow++;
            }
        }
    }
    if (win == true) {
        notifyPlayerWin()
    } else {
        notifyPlayerLoss()
    }
}

document.addEventListener('DOMContentLoaded', function() {
    gameHandler();
});

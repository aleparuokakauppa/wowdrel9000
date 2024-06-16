import GuessLetterStack from '/src/js/guessLetterStack.js';
import { handleKeyEvent } from '/src/js/keypressUtils.js';
import { drawTileChars, drawTileColors } from '/src/js/tileUtils.js';
import { checkRealWord, checkWin } from '/src/js/apiClients.js'
import { smallNotify, notifyPlayerWin, notifyPlayerLoss } from '/src/js/notify.js';
import { drawLetterColors } from '/src/js/keypadUtils.js';

async function gameHandler() {
    let currRow = 1;
    let maxRows = 6;
    let win = false;
    let guessedWords = []

    var guessLetterStack = new GuessLetterStack(5)

    // Main game loop
    while (currRow <= maxRows && win === false) {
        var currRowObj = document.getElementById('row-' + currRow);
        drawTileChars(guessLetterStack, currRowObj);
        guessLetterStack = await handleKeyEvent(guessLetterStack);
        if (guessLetterStack.guess === true) {
            var result = await checkRealWord(guessLetterStack.toString());
            if (result) {
                await checkWin(guessLetterStack.toString())
                    .then(data => {
                        drawTileColors(data, currRowObj);
                        drawLetterColors(data);
                        win = data.win;
                    })
                    .catch(error => {
                        console.error('Error in checkWin:', error);
                    })
                currRow++;
                guessLetterStack.clear();
            } else {
                guessLetterStack = new GuessLetterStack(5)
                drawTileChars(guessLetterStack, currRowObj)
                smallNotify("Not a real word.")
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

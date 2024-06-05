class GuessLetterStack {
    constructor(maxSize) {
        this.letters = new Array(maxSize).fill(' ')
        this.maxSize = maxSize;
        this.frontIndex = 0;
        this.backIndex = 0;

        // This is set to true when the
        // user wants to guess
        this.guess = false;
    }

    insertChar(item) {
        if (this.getAllocatedSize() >= this.maxSize) { return }

        this.letters[this.backIndex] = item.toUpperCase()
        this.backIndex++
    }

    removeChar() {
        if (this.getAllocatedSize() <= 0) { return }
        this.backIndex--;
        this.letters[this.backIndex] = ' ';
    }

    clear() {
        for (let index = 0; index < this.maxSize; index++) {
            this.letters[index] = ' ';
            this.guess = false;
            this.frontIndex = 0;
            this.backIndex = 0;
        }
    }

    toString() {
        return this.letters.join(" ");
    }

    toArray() {
        return this.letters;
    }

    getAllocatedSize() {
        return this.backIndex - this.frontIndex;
    }
}

function getKeyPress() {
    return new Promise((resolve) => {
        function handleKeyPress(event) {
            regex = /^[a-zA-Z]+$/;
            if (regex.test(event.key) && event.key.length === 1 
                || event.key === 'Backspace' 
                || event.key === 'Enter') {
                document.removeEventListener('keydown', handleKeyPress);
                resolve(event)
            }
        }

        document.addEventListener('keydown', handleKeyPress);
    });
}

async function handleKeyEvent(keyHistory) {
    const keyEvent = await getKeyPress();
    if (keyEvent.key === 'Backspace') {
        keyHistory.removeChar();
    }
    else if (keyEvent.key === 'Enter') {
        if (keyHistory.getAllocatedSize() === 5) {
            keyHistory.guess = true;
        }
    }
    else {
        keyHistory.insertChar(keyEvent.key);
    }
    return keyHistory
}

function drawTiles(keyHistory, tileRow) {
    let chars = keyHistory.toArray();
    let tiles = tileRow.children;
    for (let index = 0; index < chars.length; index++) {
        tiles[index].textContent = chars[index];
    }
}

function checkWin(keyHistory) {

}

async function gameHandler() {
    let currRow = 1;
    let maxRows = 6;

    var keyHistory = new GuessLetterStack(5)

    // Main game loop
    while (currRow <= maxRows) {
        keyHistory = await handleKeyEvent(keyHistory);
        if (keyHistory.guess === true) {
            // Check if word is a real word
            // Send the request for checking the correct answer
            currRow++;
            keyHistory.clear();
        } else {
            var currRowObj = document.getElementById('row-' + currRow);
            drawTiles(keyHistory, currRowObj);
        }
    }

    var popup = document.getElementById('popup');
    popup.textContent = "Game over!"
    popup.style.display = 'block';
    setTimeout(() => {
        popup.style.display = 'none';
    }, 3000) // Hide popup after 3 seconds
}

document.addEventListener('DOMContentLoaded', function() {
    gameHandler();
});

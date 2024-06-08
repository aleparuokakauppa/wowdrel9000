class GuessLetterStack {
    constructor(maxSize) {
        this.letters = new Array(maxSize).fill(' ')
        this.maxSize = maxSize;
        this.frontIndex = 0;
        this.backIndex = 0;
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
        return this.letters.join('');
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
        keyHistory.insertChar(keyEvent.key.toUpperCase());
    }
    return keyHistory
}

function drawTileChars(keyHistory, tileRow) {
    let chars = keyHistory.toArray();
    let tiles = tileRow.children;
    for (let index = 0; index < chars.length; index++) {
        tiles[index].textContent = chars[index];
    }
}

function drawTileColors(serverResponse, tileRow) {
    let tiles = tileRow.children;
    serverResponse.letters.forEach((responseLetter, responseIndex) => {
        tiles[responseIndex].classList.remove('grid-item-blank');
        switch(responseLetter.status) {
            case "match":
                tiles[responseIndex].classList.add('grid-item-green');
                break;
            case "close":
                tiles[responseIndex].classList.add('grid-item-yellow');
                break;
            case "miss":
                tiles[responseIndex].classList.add('grid-item-gray');
                break;
        }
    })
}

async function checkWin(guessWord) {
    const requestData = {
        version: 1,
        guess: guessWord // This is a string
    }
    return fetch('http://localhost:8080/guess', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestData)
    })
    .then(response => {
        return response.json()
    })
    .then(data => {
        return data
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Error: ' + error);
        throw error;
    })
}

function notifyPlayer(message) {
    var popup = document.getElementById('popup');
    popup.textContent = message
    popup.style.display = 'block';
    setTimeout(() => {
        popup.style.display = 'none';
    }, 3000) // Hide popup after 3 seconds
}

async function gameHandler() {
    let currRow = 1;
    let maxRows = 6;
    let win = false;

    var guessLetterStack = new GuessLetterStack(5)

    // Main game loop
    while (currRow <= maxRows && win === false) {
        console.log("state of win at start: ", win)
        guessLetterStack = await handleKeyEvent(guessLetterStack);
        var currRowObj = document.getElementById('row-' + currRow);
        drawTileChars(guessLetterStack, currRowObj);
        if (guessLetterStack.guess === true) {
            // Check if word is a real word
            // TODO don't call this as an async function...
            checkWin(guessLetterStack.toString())
            .then(data => {
                drawTileColors(data, currRowObj);
                win = data.win;
                console.log("win set to: ", data.win)
            })
            .catch(error => {
                console.error('Error in checkWin:', error);
            })
            currRow++;
            guessLetterStack.clear();
        }
        console.log("state of win at end: ", win)
    }
    if (win == true) {
        notifyPlayer("You won!")
    } else {
        notifyPlayer("You lost!")
    }
}

document.addEventListener('DOMContentLoaded', function() {
    gameHandler();
});

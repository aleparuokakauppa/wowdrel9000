export function drawTileChars(keyHistory, tileRow) {
    let chars = keyHistory.toArray();
    let tiles = tileRow.children;
    for (let index = 0; index < chars.length; index++) {
        tiles[index].textContent = chars[index];
    }
}

export function drawTileColors(serverResponse, tileRow) {
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

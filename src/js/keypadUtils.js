export function drawLetterColors(serverResponse) {
    serverResponse.letters.forEach((responseLetter) => {
        let letterContainer = document.getElementById('key-' + String.fromCharCode(responseLetter.char))
        switch(responseLetter.status) {
            case "match":
                letterContainer.classList.remove('key-item-blank');
                letterContainer.classList.remove('key-item-yellow');
                letterContainer.classList.add('key-item-green');
                break;
            case "close":
                letterContainer.classList.remove('key-item-blank');
                letterContainer.classList.add('key-item-yellow');
                break;
            case "miss":
                letterContainer.classList.remove('key-item-blank');
                letterContainer.classList.add('key-item-gray');
                break;
        }
    })
}

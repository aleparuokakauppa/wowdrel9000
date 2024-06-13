function getKeyPress() {
    return new Promise((resolve) => {
        function handleKeyPress(event) {
            let regex = /^[a-zA-Z]+$/;
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

export async function handleKeyEvent(keyHistory) {
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

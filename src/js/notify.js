export function smallNotify(message) {
    let popupElement = document.getElementById('small-popup');
    popupElement.textContent = message;
    popupElement.style.display = 'flex';
    setTimeout(() => {
        popupElement.style.display = 'none';
    }, 2000)
}

export function notifyPlayer(message) {
    document.getElementById('popup-message').textContent = message;
    document.getElementById('popup').style.display = 'flex';
}

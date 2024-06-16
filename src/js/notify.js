export function smallNotify(message) {
    let popupElement = document.getElementById('small-popup');
    popupElement.textContent = message;
    popupElement.style.display = 'flex';
    setTimeout(() => {
        popupElement.style.display = 'none';
    }, 2000)
}

function notifyPlayer(message, color) {
    let popupMessageElement = document.getElementById('popup-message')
    popupMessageElement.textContent = message;
    console.log("Setting bg col to: " + color)
    popupMessageElement.style.color = color
    document.getElementById('popup').style.display = 'flex';
}

export function notifyPlayerWin() {
    let message = "You win!"
    let color = "#0f0"
    notifyPlayer(message, color)
}

export function notifyPlayerLoss() {
    let message = "You lost!"
    let color = "#f00"
    notifyPlayer(message, color)
}

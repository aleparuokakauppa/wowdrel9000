export async function checkRealWord(guessWord) {
    const requestData = {
        version: 1,
        guess: guessWord
    }
    try {
        const response = await fetch('/realWord', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });
        if (!response.ok) {
            throw new Error("Fetch response wasn't ok")
        }
        let data = await response.json();
        return data.isreal;

    } catch (error) {
        console.error('Error:', error);
        alert('Error: ' + error);
        throw error;
    }
}

export async function checkWin(guessWord) {
    const requestData = {
        version: 1,
        guess: guessWord
    }
    try {
        const response = await fetch('/guess', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });
        
        if (!response.ok) {
            throw new Error("Fetch response wasn't ok")
        }

        return await response.json();
    } catch (error) {
        console.error('Error:', error);
        alert('Error: ' + error);
        throw error;
    }
}

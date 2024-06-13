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

export default GuessLetterStack

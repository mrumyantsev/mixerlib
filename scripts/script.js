const passwordLength = 55;
const isUniqueChars = true;

class PasswordGenerator {
    #charStore = {};
    #passLength = 20;
    #isUniqueChars = false;

    generate(passLength, isUniqueChars) {
        if (passLength)
            this.#passLength = passLength;
        
        if (isUniqueChars)
            this.#isUniqueChars = isUniqueChars;
        
        this.#fillArrays();
        this.#shuffleArrays();

        return this.#getRandomPassword();
    }

    #getRandomPassword() {
        const charStore = this.#charStore;

        let password = [];
        let previousKey = "";
        let currentKey;
        let ch;
        let endPasswordLength = this.#passLength;
        
        while (password.length < endPasswordLength) {
            currentKey = this.#getRandomKey();

            if (currentKey === previousKey)
                continue;
            
            ch = this.#getRandomCharFromArray(charStore[currentKey]);

            if (ch === "-")
                continue;
            
            password.push(ch);

            previousKey = currentKey;
        }

        return password.join("");
    }

    #getRandomKey() {
        const charStore = this.#charStore;
        const keys = Object.keys(charStore);
        const i = Math.floor(keys.length * Math.random());

        return keys[i];
    }

    #getRandomCharFromArray(arr) {
        let i, ch;

        while (true) {
            i = Math.floor(arr.length * Math.random());

            ch = arr[i];

            if (this.#isUniqueChars)
                arr[i] = "-";
            
            return ch;
        }
    }

    #shuffleArrays() {
        const charStore = this.#charStore;
        const keys = Object.keys(charStore);

        const forEachHandler = key =>
            charStore[key] = this.#getShuffledArray(charStore[key]);

        keys.forEach(forEachHandler);
    }

    #getShuffledArray(arr) {
        const res = [];

        let i, ch;
        
        while (res.length < arr.length) {
            i = Math.floor(arr.length * Math.random());

            ch = arr[i];

            if (ch === "-")
                continue;
            
            arr[i] = "-";
            res.push(ch);
        }

        return res;
    }

    #fillArrays() {
        const charStore = {
            symbols: this.#getFilledArray("@#$%&"),
            digits: this.#getFilledArray("0-9"),
            highLetters: this.#getFilledArray("A-Z"),
            lowLetters: this.#getFilledArray("a-z")
        }

        this.#charStore = charStore;
    }

    #getFilledArray(content) {
        if (!content)
            return;
        
        if (content.length === 3)
            if (content[1] === "-") {
                const startChar = content[0];
                const endChar = content[2];

                return this.#getCharRangeArray(startChar, endChar);
            }
        
        return [...content];
    }

    #getCharRangeArray(startChar, endChar) {
        const arr = [];
        let ch, code;

        for (ch = startChar; ch <= endChar;) {
            arr.push(ch);

            code = ch.charCodeAt();
            code++;
            ch = String.fromCharCode(code);
        }

        return arr;
    }
}

const main = () => {
    const body = document.getElementById("password-container");
    const bodyElement = $("#password-container");

    const passgen = new PasswordGenerator();

    const keydownHandler = event => {
        // If the Enter key is pressed, add new password.
        if (event.keyCode === 13)
            body.innerHTML += passgen.generate(passwordLength, isUniqueChars) + "<br>";
    }

    bodyElement.keydown(keydownHandler);

    // Add new password to password container at the beginning.
    body.innerHTML = passgen.generate(passwordLength, isUniqueChars) + "<br>";
}

window.onload = main;

// Function to get cookie value by name
function getCookie(name) {
    var value = "; " + document.cookie;
    var parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
}

// Check if the user is logged in, redirect to /login if not
const sessionCookie = getCookie("session");
if (!sessionCookie) {
    window.location.href = "/login";
}

// Get references to HTML elements
const selectDeckDropdown = document.getElementById('selectDeck');
const startButton = document.getElementById('startButton');
const deckSlot = document.getElementById('deck-slot');

// Function to fetch user's decks and populate the dropdown
function populateDeckDropdown() {
    fetch('/get-user-decks')
        .then(response => response.json())
        .then(decks => {
            decks.forEach(deck => {
                const option = document.createElement('option');
                option.value = deck;
                option.textContent = deck;
                selectDeckDropdown.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error fetching user decks:', error);
        });
}

// Populate the deck dropdown on page load
window.onload = function() {
    populateDeckDropdown();
};

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
}

function clearDeckSlot() {
    deckSlot.innerHTML = '<p class="slot-text">Deck</p>';
}

function resetDeckSlot() {
    // Reset deck slot
    deckSlot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
}

function resetHandSlot(i) {
    const handSlot = document.querySelector(`.hand-slot:nth-child(${i})`);
    handSlot.innerHTML = `<p class="slot-text">Hand</p>`;
}

function resetHandSlots() {
// Reset hand slots
    for (let i = 1; i <= 6; i++) {
        resetHandSlot(i);
    }
}


function resetSlots() {
    // Reset life points display and input
    const lifePointsDisplay = document.getElementById('lifePointsDisplay');
    lifePointsDisplay.textContent = `Life Points: 8000`;
    const lifePointsInput = document.getElementById('lifePointsInput');
    lifePointsInput.value = '';

    resetHandSlots();
    resetDeckSlot();
}



function initLifePoints(deckFileNames) {
    const dropdownContainer = document.getElementById('selectDeck').parentNode;
    dropdownContainer.innerHTML = `
    <div class="life-points">
        <span id="lifePointsDisplay">Life Points: 8000</span>
        <input type="number" id="lifePointsInput" value="" min="0">
        <button id="increaseLifePoints">+</button>
        <button id="decreaseLifePoints">-</button>
        <button id="resetDuel">Reset</button>
    </div>
    `;

    // Get references to the elements
    const lifePointsDisplay = document.getElementById('lifePointsDisplay');
    const lifePointsInput = document.getElementById('lifePointsInput');
    const increaseLifePointsButton = document.getElementById('increaseLifePoints');
    const decreaseLifePointsButton = document.getElementById('decreaseLifePoints');
    const resetDuelButton = document.getElementById('resetDuel'); // New button reference

    // Initialize the life points variable
    let lifePoints = 8000;

    // Update the display with the initial value
    lifePointsDisplay.textContent = `Life Points: ${lifePoints}`;

    // Event listener for the "+" button
    increaseLifePointsButton.addEventListener('click', () => {
        const amount = parseInt(lifePointsInput.value);
        lifePoints += amount;
        lifePointsDisplay.textContent = `Life Points: ${lifePoints}`;
        lifePointsInput.value = '';
    });

    // Event listener for the "-" button
    decreaseLifePointsButton.addEventListener('click', () => {
        const amount = parseInt(lifePointsInput.value);
        lifePoints -= amount;
        if (lifePoints < 0) {
            lifePoints = 0;
        }
        lifePointsDisplay.textContent = `Life Points: ${lifePoints}`;
        lifePointsInput.value = '';
    });

    // Event listener for the "Reset" button
    resetDuelButton.addEventListener('click', () => {
        resetDuel(deckFileNames)
    });
}

function resetDuel(deckFileNames) {
    const confirmReset = confirm("Are you sure you want to reset the duel and shuffle your cards?");
        if (confirmReset) {
            // Reset all slots and shuffle the deck
            resetSlots();
            shuffleArray(deckFileNames); // Make sure to provide the deckFileNames array
            currentIndex = 0;
            deckIsEmpty = false;
        }
}

var currentIndex = 0;
let deckIsEmpty = false; // Track whether the deck is empty
function isHandFull() {
    return currentIndex >= 6; // 6 cards in hand
}

function clickDeck(deckFileNames) {
    if (!deckIsEmpty && !isHandFull()) {
        if (currentIndex < deckFileNames.length) {
            // Load the next card image into the hand slot
            const cardFileName = deckFileNames[currentIndex];
            const handSlot = document.querySelector(`.hand-slot:nth-child(${currentIndex + 1})`);
            handSlot.innerHTML = `<img class="card-image" src="/card-images/${cardFileName}">`;
            currentIndex++;
        }

        if (currentIndex >= deckFileNames.length) {
            // Deck is empty, reset the deck slot and prevent further actions
            clearDeckSlot()
            deckIsEmpty = true;
        }
    } else if (deckIsEmpty) {
        // Deck is empty, prevent further actions
        alert('Deck is empty');
    } else {
        // Hand is full, display a popup
        alert('Hand is full');
    }
}

// Function to start the duel
function startDuel(deckFileNames) {
    // Replace the dropdown with the life points display and controls
    shuffleArray(deckFileNames);
    initLifePoints(deckFileNames);

    currentIndex = 0;


    // Event listener for clicking the deck
    deckSlot.addEventListener('click', () => {
        clickDeck(deckFileNames);
    });
}

// Event listener for the "Start!" button
startButton.addEventListener('click', () => {
    // Fetch the selected deck's card file names using the /cards-in-deck route
    const selectedDeck = selectDeckDropdown.value;
    fetch(`/cards-in-deck?user_id=${sessionCookie}&deck=${selectedDeck}`)
        .then(response => response.json())
        .then(deckFileNames => {
            // Display the card-back image in the hand slot 7
            deckSlot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
            // Start the duel simulation
            startDuel(deckFileNames);
        })
        .catch(error => {
            console.error('Error fetching card file names:', error);
        });
});
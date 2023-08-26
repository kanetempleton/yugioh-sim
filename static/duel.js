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

// Function to start the duel
function startDuel() {
    // Replace the dropdown with the life points display and controls
    const dropdownContainer = document.getElementById('selectDeck').parentNode;
    dropdownContainer.innerHTML = `
    <div class="life-points">
        <span id="lifePointsDisplay">Life Points: 8000</span>
        <input type="number" id="lifePointsInput" value="" min="0">
        <button id="increaseLifePoints">+</button>
        <button id="decreaseLifePoints">-</button>
    </div>
    `;

    // Get references to the elements
    const lifePointsDisplay = document.getElementById('lifePointsDisplay');
    const lifePointsInput = document.getElementById('lifePointsInput');
    const increaseLifePointsButton = document.getElementById('increaseLifePoints');
    const decreaseLifePointsButton = document.getElementById('decreaseLifePoints');

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
}


// Event listener for the "Start!" button
startButton.addEventListener('click', () => {
    // Display the card-back image in the hand slot 7
    deckSlot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
    // You can add more functionality here to start the duel
    startDuel();
});
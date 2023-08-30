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
const selectExtraDeckDropdown = document.getElementById('selectExtraDeck')
const startButton = document.getElementById('startButton');
const deckSlot = document.getElementById('slot-14');


const slotIDs = {
    DECK_SLOT: 14,
    GRAVEYARD_SLOT: 7,
    FIELD_SLOT: 1,
    SPELL_TRAP_SLOT_1: 13,
    SPELL_TRAP_SLOT_2: 12,
    SPELL_TRAP_SLOT_3: 11,
    SPELL_TRAP_SLOT_4: 10,
    SPELL_TRAP_SLOT_5: 9,
    EXTRA_DECK_SLOT: 8,
    MONSTER_SLOT_1: 6,
    MONSTER_SLOT_2: 5,
    MONSTER_SLOT_3: 4,
    MONSTER_SLOT_4: 3,
    MONSTER_SLOT_5: 2,
    HAND_SLOT_1: 15,
    HAND_SLOT_2: 16,
    HAND_SLOT_3: 17,
    HAND_SLOT_4: 18,
    HAND_SLOT_5: 19,
    HAND_SLOT_6: 20,
    HAND_SLOT_7: 21,
};

function getSlotID(slot) {
    const slotIdHtml = slot.getAttribute('id');
    switch (slotIdHtml) {
        case "slot-14":
            return slotIDs.DECK_SLOT;
        case "slot-1":
            return slotIDs.FIELD_SLOT;
        case "slot-8":
            return slotIDs.EXTRA_DECK_SLOT;
        case "slot-7":
            return slotIDs.GRAVEYARD_SLOT;
        case "slot-15":
            return slotIDs.HAND_SLOT_1;
        case "slot-16":
            return slotIDs.HAND_SLOT_2;
        case "slot-17":
            return slotIDs.HAND_SLOT_3;
        case "slot-18":
            return slotIDs.HAND_SLOT_4;
        case "slot-19":
            return slotIDs.HAND_SLOT_5;
        case "slot-20":
            return slotIDs.HAND_SLOT_6;
        case "slot-21":
            return slotIDs.HAND_SLOT_7;
        case "slot-6":
            return slotIDs.MONSTER_SLOT_1;
        case "slot-5":
            return slotIDs.MONSTER_SLOT_2;
        case "slot-4":
            return slotIDs.MONSTER_SLOT_3;
        case "slot-3":
            return slotIDs.MONSTER_SLOT_4;
        case "slot-2":
            return slotIDs.MONSTER_SLOT_5;
        case "slot-13":
            return slotIDs.SPELL_TRAP_SLOT_1;
        case "slot-12":
            return slotIDs.SPELL_TRAP_SLOT_2;
        case "slot-11":
            return slotIDs.SPELL_TRAP_SLOT_3;
        case "slot-10":
            return slotIDs.SPELL_TRAP_SLOT_4;
        case "slot-9":
            return slotIDs.SPELL_TRAP_SLOT_5;
    }
    return null
}
function getSlot(slotID) {
    switch (slotID) {
        case slotIDs.DECK_SLOT:
            return document.getElementById('slot-14');
        case slotIDs.GRAVEYARD_SLOT:
            return document.getElementById('slot-7');
        case slotIDs.FIELD_SLOT:
            return document.getElementById('slot-1');
        case slotIDs.EXTRA_DECK_SLOT:
            return document.getElementById('slot-8');
        case slotIDs.SPELL_TRAP_SLOT_1:
            return document.getElementById('slot-13');
        case slotIDs.SPELL_TRAP_SLOT_2:
            return document.getElementById('slot-12');
        case slotIDs.SPELL_TRAP_SLOT_3:
            return document.getElementById('slot-11');
        case slotIDs.SPELL_TRAP_SLOT_4:
            return document.getElementById('slot-10');
        case slotIDs.SPELL_TRAP_SLOT_5:
            return document.getElementById('slot-9');
            case slotIDs.MONSTER_SLOT_1:
                return document.getElementById('slot-6');
            case slotIDs.MONSTER_SLOT_2:
                return document.getElementById('slot-5');
            case slotIDs.MONSTER_SLOT_3:
                return document.getElementById('slot-4');
            case slotIDs.MONSTER_SLOT_4:
                return document.getElementById('slot-3');
            case slotIDs.MONSTER_SLOT_5:
                return document.getElementById('slot-2');
        case slotIDs.HAND_SLOT_1:
            return document.getElementById('slot-15')
            case slotIDs.HAND_SLOT_2:
                return document.getElementById('slot-16');
            case slotIDs.HAND_SLOT_3:
                return document.getElementById('slot-17');
            case slotIDs.HAND_SLOT_4:
                return document.getElementById('slot-18');
            case slotIDs.HAND_SLOT_5:
                return document.getElementById('slot-19');
            case slotIDs.HAND_SLOT_6:
                return document.getElementById('slot-20');
            case slotIDs.HAND_SLOT_7:
                return document.getElementById('slot-21');
        default:
            return null; 
    }
}

const graveyardSlot = getSlot(slotIDs.GRAVEYARD_SLOT);
const fieldSlot = getSlot(slotIDs.FIELD_SLOT);
const extraDeckSlot = getSlot(slotIDs.EXTRA_DECK_SLOT)
function handSlot(n) { 
    if (n > 7 || n < 1)
        return 
    return getSlot(n+14)
}
function monsterSlot(n) { 
    if (n > 5 || n < 1)
        return 
    return getSlot(14-n)
}
function spellTrapSlot(n) { 
    if (n > 5 || n < 1)
        return 
    return getSlot(7-n)
}



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
                selectExtraDeckDropdown.appendChild(option);
            });
        })
        .catch(error => {
            console.error('Error fetching user decks:', error);
        });
}

function populateExtraDeckDropdown() {
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
    populateExtraDeckDropdown();
};

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
}

function resetSlotID(i) {
    console.log("reset slot ID: "+i)
    makeVertical(getSlot(i))
    switch (i) {
        case slotIDs.DECK_SLOT:
            resetDeckSlot();
            break;
        case slotIDs.EXTRA_DECK_SLOT:
            resetExtraDeckSlot();
            break;
        case slotIDs.HAND_SLOT_1:
            resetHandSlot(1);
            break;
        case slotIDs.HAND_SLOT_2:
            resetHandSlot(2);
            break;
        case slotIDs.HAND_SLOT_3:
            resetHandSlot(3);
            break;
        case slotIDs.HAND_SLOT_4:
            resetHandSlot(4);
            break;
        case slotIDs.HAND_SLOT_5:
            resetHandSlot(5);
            break;
        case slotIDs.HAND_SLOT_6:
            resetHandSlot(6);
            break;
        case slotIDs.HAND_SLOT_7:
            resetHandSlot(7);
            break;
        case slotIDs.FIELD_SLOT:
            resetFieldSlot();
            break;
        case slotIDs.GRAVEYARD_SLOT:
            resetGraveyardSlot();
            break;
        case slotIDs.MONSTER_SLOT_1:
            resetMonsterSlot(1);
            break;
        case slotIDs.MONSTER_SLOT_2:
            resetMonsterSlot(2);
            break;
        case slotIDs.MONSTER_SLOT_3:
            resetMonsterSlot(3);
            break;
        case slotIDs.MONSTER_SLOT_4:
            resetMonsterSlot(4);
            break;
        case slotIDs.MONSTER_SLOT_5:
            resetMonsterSlot(5);
            break;
        case slotIDs.SPELL_TRAP_SLOT_1:
            resetSpellTrapSlot(1);
            break;
        case slotIDs.SPELL_TRAP_SLOT_2:
            resetSpellTrapSlot(2);
            break;
        case slotIDs.SPELL_TRAP_SLOT_3:
            resetSpellTrapSlot(3);
            break;
        case slotIDs.SPELL_TRAP_SLOT_4:
            resetSpellTrapSlot(4);
            break;
        case slotIDs.SPELL_TRAP_SLOT_5:
            resetSpellTrapSlot(5);
            break;

    }
}
function resetGraveyardSlot() {
    console.log("resetGraveyardSlot")
    graveyardSlot.innerHTML = '<p class="slot-text">Graveyard</p>'
}

function resetFieldSlot() {
    console.log("resetFieldSlot")
    fieldSlot.innerHTML = '<p class="slot-text">Field Card</p>'
}

function setFacedownCard(slot) {
    console.log("setFacedownCard: "+getSlotID(slot))
    slot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
}

function resetExtraDeckSlot() {
    console.log("resetExtraDeckSlot")
    extraDeckSlot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
}

function clearExtraDeckSlot() {
    console.log("clearExtraDeckSlot")
    extraDeckSlot.innerHTML = '<p class="slot-text">Extra Deck</p>'
}

function clearDeckSlot() {
    console.log("clearDeckSlot")
    deckSlot.innerHTML = '<p class="slot-text">Deck</p>';
}

function resetDeckSlot() {
    console.log("resetDeckSlot")
    // Reset deck slot
    deckSlot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
}

function resetHandSlot(i) {
    const handSlot = getSlot(14+i);
    handSlot.innerHTML = `<p class="slot-text">Hand</p>`;
}

function resetMonsterSlot(i) {
    console.log("resetMonsterSlot: "+i)
    if (i>5 || i < 1)
        return 
    getSlot(7-i).innerHTML = '<p class="slot-text">Monster</p>';
}
function resetSpellTrapSlot(i) {
    console.log("resetSpellTrapSlot")
    if (i>5 || i < 1)
        return 
    getSlot(14-i).innerHTML = '<p class="slot-text">Spell/Trap</p>';
}

function resetMonsterSlots() {
    console.log("resetMonsterSlots")
    for (let i=1;i<=5;i++)
        resetMonsterSlot(i)
}
function resetSpellTrapSlots() {
    console.log("resetSpellTrapSlots")
    for (let i=1;i<=5;i++)
        resetSpellTrapSlot(i)
}

function resetHandSlots() {
    console.log("resetHandSlots")
// Reset hand slots
    for (let i = 1; i <= 6; i++) {
        resetHandSlot(i);
    }
}


function resetSlots() {
    console.log("resetSlots")
    // Reset life points display and input
    const lifePointsDisplay = document.getElementById('lifePointsDisplay');
    lifePointsDisplay.textContent = `Life Points: 8000`;
    const lifePointsInput = document.getElementById('lifePointsInput');
    lifePointsInput.value = '';

    resetHandSlots();
    resetDeckSlot();
    resetMonsterSlots();
    resetSpellTrapSlots();

}

function createOverlayComponent() {
    const overlay = document.createElement('div');
    overlay.classList.add('card-overlay');
    const overlayContent = document.createElement('div');
    overlayContent.classList.add('overlay-buttons');
    /*overlayContent.innerHTML = `
        <button class="overlay-button overlay-button-view">View</button>
        <button class="overlay-button overlay-button-play">Play</button>
    `;*/
    overlay.appendChild(overlayContent);
    return overlay;
}

function addOverlay(cardImage, cardFileName) {
    const overlay = document.createElement('div');
    overlay.classList.add('card-overlay');
    const overlayContent = document.createElement('div');
    overlayContent.classList.add('overlay-buttons');
    overlayContent.innerHTML = `
        <button class="overlay-button view-button">View</button>
        <button class="overlay-button play-button">Play</button>
    `;
    overlay.appendChild(overlayContent);

    cardImage.parentElement.appendChild(overlay);

    cardImage.addEventListener('mouseleave', () => {
        overlay.remove();
    });
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
        resetDuel(deckFileNames,extraDeckFileNames)
    });
}

function resetDuel(deckFileNames,extraDeckFileNames) {
    const confirmReset = confirm("Are you sure you want to reset the duel and shuffle your cards?");
        if (confirmReset) {
            // Reset all slots and shuffle the deck
            resetSlots();
            shuffleArray(deckFileNames); // Make sure to provide the deckFileNames array
            currentIndex = 0;
            deckIsEmpty = false;
        }
}
function getCardImageNameFromSlotID(slotID) {
    console.log("getCardImageNameFromSlotID: "+slotID)
    const slotElement = getSlot(slotID);
    if (slotElement) {
        console.log("found slot element")
        const cardImageElement = slotElement.querySelector('img');
        if (cardImageElement) {
            console.log("found card image element")
            const imagePath = cardImageElement.getAttribute('src');
            const imageName = imagePath.split('/').pop();
            return imageName;
        }
    }
    return null; // Return null if no card image is found in the slot
}
function showGraveyardCards() {
    // Create the overlay container
    const overlayContainer = document.createElement('div');
    overlayContainer.classList.add('graveyard-overlay');

    // Create the close button
    const closeButton = document.createElement('button');
    closeButton.textContent = 'Close';
    closeButton.classList.add('overlay-close-button');
    closeButton.addEventListener('click', () => {
        overlayContainer.remove();
    });

    // Create a grid to display the card images
    const gridContainer = document.createElement('div');
    gridContainer.classList.add('grid-container');

    // Function to rebuild the grid with updated card images
    const rebuildGrid = () => {
        gridContainer.innerHTML = ''; // Clear existing images
        graveyardCardFiles.forEach(cardFileName => {
            const cardImage = document.createElement('img');
            cardImage.src = `/card-images/${cardFileName}`;
            cardImage.classList.add('graveyard-card-image');
            
            // Add click event to call returnToHand method
            cardImage.addEventListener('click', () => {
                returnToHand(cardFileName); // Call your method here
                rebuildGrid(); // Refresh the grid
            });

            gridContainer.appendChild(cardImage);
        });
    };

    // Initial build of the grid
    rebuildGrid();

    // Append elements to the overlay container
    overlayContainer.appendChild(closeButton);
    overlayContainer.appendChild(gridContainer);

    // Append the overlay container to the body
    document.body.appendChild(overlayContainer);
}




let graveyardCardFiles = null;
function addCardToGraveyard(cardFileName) {
    console.log("addCardToGraveyard: " + cardFileName);
    
    if (graveyardCardFiles === null) {
        graveyardCardFiles = []; // Initialize the list if it's null
    }
    
    graveyardCardFiles.push(cardFileName); // Add the card file name to the list
}



function addCardToSlot(slot, cardFileName, vertical) {
    
    console.log("addCardToSlot: "+getSlotID(slot)+","+cardFileName)
    

    if (slotID(slot) == slotIDs.GRAVEYARD_SLOT) {
        addCardToGraveyard(cardFileName);
        return;
    }
    slot.innerHTML = `<img class="card-image" src="/card-images/${cardFileName}">`;
    if (vertical)
        slot.classList.add('vertical-card');
    else
        slot.classList.add('horizontal-card')
}

function moveCardFromSlot(slot_from, slot_to, facedown) {
    console.log("moveCardFromSlot: "+getSlotID(slot_from)+","+getSlotID(slot_to)+","+facedown)
    if (facedown) {
        setFacedownCard(slot_to);
        if (isHorizontal(slot_from)) {
            makeHorizontal(slot_to);
        }
        resetSlotID(getSlotID(slot_from));
    } else {
        
        card = getCardImageNameFromSlotID(getSlotID(slot_from));
        addCardToSlot(slot_to,card,isVertical(slot_from));
        if (isHorizontal(slot_from)) {
            makeHorizontal(slot_to);
        }
        resetSlotID(getSlotID(slot_from));
    }
    setSlotCard(slot_to,cardInSlot[slotID(slot_from)])
    clearSlotCard(slot_from)
    unselectCard(slot_from)
}

function drawFromDeck() {
    const cardFileName = cardFileNames[currentIndex];
    const handSlot = firstOpenHandSlot();
    handSlot.innerHTML = `<img class="card-image" src="/card-images/${cardFileName}">`;
    makeVertical(handSlot)
    setSlotCard(handSlot,cardFileName);
    currentIndex++;
}

function returnToHand(cardFileName) {
    console.log("returnToHand: " + cardFileName);
    if (isHandFull()) {
        // do nothing
    } else {
        const handSlot = firstOpenHandSlot();
        handSlot.innerHTML = `<img class="card-image" src="/card-images/${cardFileName}">`;
        makeVertical(handSlot);
        setSlotCard(handSlot, cardFileName);

        // Find the index of the cardFileName in graveyardCardFiles
        const index = graveyardCardFiles.indexOf(cardFileName);
        
        // Remove the cardFileName from the array if found
        if (index !== -1) {
            graveyardCardFiles.splice(index, 1);
        }
    }
}

function addExtraDeckCardToHand(cardFileName) {
    console.log("addExtraDeckCardToHand: " + cardFileName);
    if (isHandFull()) {
        // do nothing
    } else {
        const handSlot = firstOpenHandSlot();
        handSlot.innerHTML = `<img class="card-image" src="/card-images/${cardFileName}">`;
        makeVertical(handSlot);
        setSlotCard(handSlot, cardFileName);

        // Find the index of the cardFileName in graveyardCardFiles
        const index = extraDeckCards.indexOf(cardFileName);
        
        // Remove the cardFileName from the array if found
        if (index !== -1) {
            extraDeckCards.splice(index, 1);
        }
    }
}

function clickExtraDeck() {
    // Create the overlay container
    const overlayContainer = document.createElement('div');
    overlayContainer.classList.add('graveyard-overlay');

    // Create the close button
    const closeButton = document.createElement('button');
    closeButton.textContent = 'Close';
    closeButton.classList.add('overlay-close-button');
    closeButton.addEventListener('click', () => {
        overlayContainer.remove();
    });

    // Create a grid to display the card images
    const gridContainer = document.createElement('div');
    gridContainer.classList.add('grid-container');

    // Function to rebuild the grid with updated card images
    const rebuildGrid = () => {
        gridContainer.innerHTML = ''; // Clear existing images
        extraDeckCards.forEach(cardFileName => {
            const cardImage = document.createElement('img');
            cardImage.src = `/card-images/${cardFileName}`;
            cardImage.classList.add('graveyard-card-image');

            cardImage.addEventListener('click', () => {
                addExtraDeckCardToHand(cardFileName); // Call your method here
                rebuildGrid(); // Refresh the grid
            });

            gridContainer.appendChild(cardImage);
        });
    };

    // Initial build of the grid
    rebuildGrid();

    // Append elements to the overlay container
    overlayContainer.appendChild(closeButton);
    overlayContainer.appendChild(gridContainer);

    // Append the overlay container to the body
    document.body.appendChild(overlayContainer);
}

function clickDeck() {
    if (!isHandFull()) 
        console.log("clickDeck; first empty hand slot = "+slotID(firstOpenHandSlot()))
    else 
        console.log("clickDeck; first empty hand slot = "+firstOpenHandSlot())

    if (!deckEmpty()) {
        if (!isHandFull()) {
            drawFromDeck();
            if (deckEmpty()) { //drew the last card
                clearDeckSlot()
                deckIsEmpty = true;
            }
        } else {
            alert('Hand is full')
        }
    } else {
        // Deck is empty, prevent further actions
        alert('Deck is empty');
    }
}

var currentIndex = 0;
let deckIsEmpty = false; // Track whether the deck is empty
let selectedCardSlot = null;
let selectedCardImage = null;
let cardOverlayVisible = false;
let inSelectMode = false;

function deckEmpty() {
    return currentIndex >= cardFileNames.length;
}

function isHandFull() {
    var val = hasCard(getSlot(slotIDs.HAND_SLOT_1)) && hasCard(getSlot(slotIDs.HAND_SLOT_2))
            && hasCard(getSlot(slotIDs.HAND_SLOT_3)) && hasCard(getSlot(slotIDs.HAND_SLOT_4))
            && hasCard(getSlot(slotIDs.HAND_SLOT_5)) && hasCard(getSlot(slotIDs.HAND_SLOT_6));
    console.log("isHandFull: "+val)
    return val;
}

function slotID(slot) {
    return getSlotID(slot);
}

function clickPlayButton(cardSlot) {
    console.log("clickPlayButton: "+slotID(cardSlot))
    if (selectedCardSlot === null) {
        selectCard(cardSlot);
    } else if (selectedCardSlot === cardSlot) {
        unselectCard(cardSlot)
    } else if (!cardSlot.querySelector('.card-image')) {
        // Move the selected card to the new slot
        cardSlot.innerHTML = '';
        cardSlot.appendChild(selectedCardImage);
        selectedCardSlot.innerHTML = '';
        selectedCardSlot.classList.remove('selected-card');
        selectedCardSlot.appendChild(cardImage);
        selectedCardSlot = null;
        selectedCardImage = null;
    }
}

var cardFileNames = null;

// Function to start the duel
function startDuel(deckFileNames,extraDeckFileNames) {
    // Replace the dropdown with the life points display and controls
    cardFileNames = deckFileNames;
    shuffleArray(cardFileNames);
    initLifePoints(cardFileNames,extraDeckFileNames);
    extraDeckCards = extraDeckFileNames;

    currentIndex = 0;

    const cardSlots = document.querySelectorAll('.grid-slot, .hand-slot');
    cardSlots.forEach(cardSlot => {

        cardSlot.addEventListener('mouseenter', () => {
            mouseEnterSlot(cardSlot)
        });

        cardSlot.addEventListener('click', () => {
            clickSlot(cardSlot)
        });
    });
}


// true if the card slot has a card image
function hasCard(cardSlot) {
    const cardImage = cardSlot.querySelector('.card-image');
    if (cardImage)
        return true;
    return false;
}

function hasSelectedCard() {
    return selectedCardSlot != null;
}

function clickGraveyard() {
    console.log("clickGraveyard")
    if (!hasSelectedCard()) {
        showGraveyardCards();
    }
}
function showDeckCards() {
    // Create the overlay container
    const overlayContainer = document.createElement('div');
    overlayContainer.classList.add('deck-overlay');

    // Create the close button
    const closeButton = document.createElement('button');
    closeButton.textContent = 'Close';
    closeButton.classList.add('overlay-close-button');
    closeButton.addEventListener('click', () => {
        overlayContainer.remove();
    });

    // Create a grid to display the card images
    const gridContainer = document.createElement('div');
    gridContainer.classList.add('grid-container');

    // Get the remaining cards in the deck (replace with your actual deck info)
    const remainingCardFiles = cardFileNames.slice(currentIndex);

    // Shuffle the array to randomize the order
    const shuffledCardFiles = shuffleArray2(remainingCardFiles);

    // Function to rebuild the grid with updated card images
    const rebuildGrid = () => {
        gridContainer.innerHTML = ''; // Clear existing images
        shuffledCardFiles.forEach(cardFileName => {
            const cardImage = document.createElement('img');
            cardImage.src = `/card-images/${cardFileName}`;
            cardImage.classList.add('deck-card-image');
            
            // Add click event to call addToHand method
            cardImage.addEventListener('click', () => {
                addToHand(cardFileName); // Call your method here
                overlayContainer.remove();
            });

            gridContainer.appendChild(cardImage);
        });
    };

    // Initial build of the grid
    rebuildGrid();

    // Append elements to the overlay container
    overlayContainer.appendChild(closeButton);
    overlayContainer.appendChild(gridContainer);

    // Append the overlay container to the body
    document.body.appendChild(overlayContainer);
}
function addToHand(cardFileName) {
    console.log("addCardFromDeck: " + cardFileName);
    if (isHandFull()) {
        // do nothing
    } else {
        const handSlot = firstOpenHandSlot();
        handSlot.innerHTML = `<img class="card-image" src="/card-images/${cardFileName}">`;
        makeVertical(handSlot);
        setSlotCard(handSlot, cardFileName);

        // Find the index of the cardFileName in cardFileNames
        const index = cardFileNames.indexOf(cardFileName);
        
        // Remove the cardFileName from the array if found
        if (index !== -1) {
            cardFileNames.splice(index, 1);
        }
    }
}


// Function to shuffle an array
function shuffleArray2(array) {
    const shuffledArray = array.slice(); // Create a copy of the array
    for (let i = shuffledArray.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [shuffledArray[i], shuffledArray[j]] = [shuffledArray[j], shuffledArray[i]]; // Swap elements
    }
    return shuffledArray;
}




// handles clicking logic for card slots
function clickSlot(slot) {
    console.log("clickSlot: "+getSlotID(slot))
    if (getSlotID(slot)==slotIDs.DECK_SLOT) {
        clickDeck()
        return;
    }
    if (getSlotID(slot)==slotIDs.GRAVEYARD_SLOT) {
        clickGraveyard();
    }
    if (getSlotID(slot)==slotIDs.EXTRA_DECK_SLOT) {
        clickExtraDeck();
    }
    if (getSlotID(slot)==slotIDs.HAND_SLOT_7) {
        showDeckCards();
    }
    if (selectedCardSlot != null) {
        if (!cardOverlayVisible) {
            if (slot == selectedCardSlot) {
                unselectCard(slot)
            } else {
                // target a slot for a card
                if (getSlotID(slot)!=slotIDs.EXTRA_DECK_SLOT && getSlotID(slot)!=slotIDs.DECK_SLOT 
                && getSlotID(slot)!=slotIDs.HAND_SLOT_7) {
                    // TODO: add card to extra deck
                
                    moveCardFromSlot(selectedCardSlot,slot,false);
                    
                } 
            }
        }
    }
}

function createActionOverlay(cardSlot) {
    const cardImage = cardSlot.querySelector('.card-image');
    const overlay = createOverlayComponent();
    
    cardSlot.appendChild(overlay);
    cardOverlayVisible = true;
    overlay.addEventListener('mouseleave', () => {
        overlay.remove();
        cardOverlayVisible = false;
    });

    // Create a container for the buttons
    const buttonsContainer = document.createElement('div');
    buttonsContainer.classList.add('buttons-container');
    overlay.appendChild(buttonsContainer);

    // Add "Play" button
    const playButton = document.createElement('button');
    playButton.classList.add('overlay-button', 'overlay-button-small', 'flip-button');
    playButton.textContent = 'Play';
    playButton.addEventListener('click', () => {
        clickPlayButton(cardSlot);
    });
    buttonsContainer.appendChild(playButton);

    // Add "View" button
    const viewButton = document.createElement('button');
    viewButton.classList.add('overlay-button', 'overlay-button-small', 'flip-button');
    viewButton.textContent = 'View';
    viewButton.addEventListener('click', () => {
        const fullCardImage = cardImage.cloneNode();
        fullCardImage.classList.add('overlay-content');

        // Create the full-page overlay
        const fullPageOverlay = document.createElement('div');
        fullPageOverlay.classList.add('overlay-full-page');
        fullPageOverlay.appendChild(fullCardImage);

        // Append the full-page overlay to the body
        document.body.appendChild(fullPageOverlay);

        // Add event listener to close the full-page overlay
        fullPageOverlay.addEventListener('click', () => {
            fullPageOverlay.remove();
        });
    });
    buttonsContainer.appendChild(viewButton);

    // Add "Flip" button
    const flipButton = document.createElement('button');
    flipButton.classList.add('overlay-button', 'overlay-button-small', 'flip-button');
    flipButton.textContent = 'Flip';
    buttonsContainer.appendChild(flipButton);

    // Add "Change Position" button
    const changePositionButton = document.createElement('button');
    changePositionButton.classList.add('overlay-button', 'overlay-button-small', 'change-position-button');
    changePositionButton.textContent = 'ATK/DEF';
    buttonsContainer.appendChild(changePositionButton);

    // Add event listeners for "Flip" and "Change Position" buttons
    flipButton.addEventListener('click', () => {
        // Call flipCard(slot) function here when implemented
        flipCard(cardSlot)
    });

    changePositionButton.addEventListener('click', () => {
        // Call changePosition(slot) function here
        changePosition(cardSlot)
    });
}
function isVertical(slot) {
    return slot.classList.contains('vertical-card')
}
function isHorizontal(slot) {
    return slot.classList.contains('horizontal-card')
}
function makeVertical(slot) {
    slot.classList.remove('horizontal-card')
    slot.classList.add('vertical-card')
}
function makeHorizontal(slot) {
    slot.classList.remove('vertical-card')
    slot.classList.add('horizontal-card')
}
function changePosition(slot) {
    const cardImage = slot.querySelector('.card-image');
    if (!cardImage) {
        console.log(`No card image found in slot ${slot}.`);
        return;
    }

    const isVertical = slot.classList.contains('vertical-card');
    
    if (isVertical) {
        // Switch to horizontal position
        slot.classList.remove('vertical-card');
        slot.classList.add('horizontal-card');
        cardImage.style.transform = 'rotate(90deg)';
    } else {
        // Switch to vertical position
        slot.classList.remove('horizontal-card');
        slot.classList.add('vertical-card');
        cardImage.style.transform = 'none';
    }
}


function isFaceDown(slot) {
    const cardImageElement = slot.querySelector('.card-image');
        if (cardImageElement) {
            const imagePath = cardImageElement.getAttribute('src');
            const val = imagePath.endsWith('card-back.jpg');
            console.log("isFaceDown: "+slotID(slot)+","+val)
            return val;
        }
        const val2 = false;
        console.log("isFaceDown: "+slotID(slot)+","+val2)
        return val2;
}

function flipCard(slot) {
    console.log("flipCard: "+slotID(slot))
    const cardImageElement = slot.querySelector('.card-image');
        if (cardImageElement) {
            const slotID = getSlotID(slot);

            if (isFaceDown(slot)) {
                console.log("Flipping face-down card at " + slotID + " to face-up");
                const cardFileName = cardInSlot[slotID];
                cardImageElement.setAttribute('src', '/card-images/'+cardFileName);
            } else {
                console.log("Flipping face-up card at " + slotID + " to face-down");
                cardImageElement.setAttribute('src', '/card-images/card-back.jpg');
            }
        }
}


function createTargetOverlay(cardSlot) {
    const emptySlotOverlay = document.createElement('div');
    emptySlotOverlay.classList.add('empty-slot-overlay');
    cardSlot.appendChild(emptySlotOverlay);
    cardSlot.addEventListener('mouseleave', () => {
        const emptySlotOverlay = cardSlot.querySelector('.empty-slot-overlay');
        if (emptySlotOverlay) {
            emptySlotOverlay.remove();
        }
    });
}

// handle mouseover functions for slots
function mouseEnterSlot(cardSlot) {
    const cardImage = cardSlot.querySelector('.card-image');
            const overlay = createOverlayComponent();
            
            if (cardImage && cardSlot !== deckSlot && cardSlot !== extraDeckSlot) {
                createActionOverlay(cardSlot)
            }

            const playButton = overlay.querySelector('.overlay-button-play');
            if (playButton) {
                playButton.addEventListener('click', () => {
                    clickPlayButton(cardSlot);
                });
            }

   

            if (!cardImage && cardSlot !== deckSlot && selectedCardSlot !== cardSlot) {
                if (selectedCardSlot != null) {
                    createTargetOverlay(cardSlot);
                }
            }
}

function firstOpenHandSlot() {
    if (!hasCard(getSlot(slotIDs.HAND_SLOT_1)))
        return getSlot(slotIDs.HAND_SLOT_1);
    if (!hasCard(getSlot(slotIDs.HAND_SLOT_2)))
        return getSlot(slotIDs.HAND_SLOT_2);
    if (!hasCard(getSlot(slotIDs.HAND_SLOT_3)))
        return getSlot(slotIDs.HAND_SLOT_3);
    if (!hasCard(getSlot(slotIDs.HAND_SLOT_4)))
        return getSlot(slotIDs.HAND_SLOT_4);
    if (!hasCard(getSlot(slotIDs.HAND_SLOT_5)))
        return getSlot(slotIDs.HAND_SLOT_5);
    if (!hasCard(getSlot(slotIDs.HAND_SLOT_6)))
        return getSlot(slotIDs.HAND_SLOT_6);
    return null;
}

function selectCard(cardSlot) {
    console.log("selecting card: "+getSlotID(cardSlot))
    const cardImage = cardSlot.querySelector('.card-image');
    selectedCardSlot = cardSlot;
    selectedCardImage = cardImage.cloneNode();
    cardImage.classList.add('selected-card');
    inSelectMode = true;
}

function unselectCard(cardSlot) {
    console.log("unselecting card "+getSlotID(cardSlot))
    const cardImage = cardSlot.querySelector('.card-image');
    inSelectMode = false;
    selectedCardSlot = null;
    selectedCardImage = null;
    if (cardImage)
        cardImage.classList.remove('selected-card');
}

var cardInSlot = {};
function setSlotCard(slot,fileName) {
    console.log("setSlotCard: "+slotID(slot)+","+fileName)
    cardInSlot[getSlotID(slot)] = fileName;
}
function clearSlotCard(slot) {
    console.log("clearSlotCard: "+slotID(slot))
    cardInSlot[getSlotID(slot)] = null;
}

var extraDeckCards = null;
// Event listener for the "Start!" button
startButton.addEventListener('click', () => {
    // Fetch the selected deck's card file names using the /cards-in-deck route
    const selectedDeck = selectDeckDropdown.value;
    const extraDeck = selectExtraDeckDropdown.value;

    fetch(`/cards-in-deck?user_id=${sessionCookie}&deck=${selectedDeck}`)
        .then(response => response.json())
        .then(deckFileNames => {
            // Display the card-back image in the hand slot 7
            deckSlot.innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;
            getSlot(slotIDs.EXTRA_DECK_SLOT).innerHTML = `<img class="card-image" src="/card-images/card-back.jpg">`;

            // Fetch the extra deck's card file names using the /cards-in-deck route
            return fetch(`/cards-in-deck?user_id=${sessionCookie}&deck=${extraDeck}`)
                .then(response => response.json())
                .then(extraDeckFileNames => {
                    // Start the duel simulation with both deck and extra deck file names
                    startDuel(deckFileNames, extraDeckFileNames);
                });
        })
        .catch(error => {
            console.error('Error fetching card file names:', error);
        });
});

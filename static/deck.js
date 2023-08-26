
// Fetch user's cards and populate the cardCanvas div
        // Function to load user's cards
        function getUrlParam(name) {
            const urlParams = new URLSearchParams(window.location.search);
            return urlParams.get(name);
        }
            function loadUserCards() {
                const deckParam = getUrlParam('deck');
                const url = deckParam ? `/get-user-cards?deck=${deckParam}` : '/get-user-cards';
                if (deckParam && deckParam !== 'None') {
                    const deckNameElement = document.createElement('p');
                    deckNameElement.textContent = `Viewing Deck: ${deckParam}`;
                    deckNameElement.style.fontWeight = 'bold';
                    deckNameElement.style.fontSize = '36px';  // Adjust the font size as needed
                    deckNameElement.style.color = '#FDF5E6';
                    document.body.insertBefore(deckNameElement, document.body.firstChild);
                }
    
                
    
                //const fetchUrl = deckParam !== 'None' ? `/get-user-cards?deck=${deckParam}` : '/get-user-cards';
    
                fetch(url)
                    .then(response => response.json())
                    .then(cardIdentifiers => {
                    const cardCanvas = document.getElementById('cardCanvas');
                    cardCanvas.innerHTML = ''; // Clear existing content
    
                    cardIdentifiers.forEach(identifier => {




                        const decodedIdentifier = decodeURIComponent(identifier);
                        const cardContainer = document.createElement('div');
                        cardContainer.classList.add('card');
    
                        const cardImage = document.createElement('img');
                        cardImage.src = `card-images/${decodedIdentifier}`;
                        cardImage.alt = decodedIdentifier;
    
                        // Hover effect: Show actions on card hover
                        cardContainer.addEventListener('mouseenter', () => {
                            cardActions.style.display = 'block';
                        });
    
                        cardContainer.addEventListener('mouseleave', () => {
                            cardActions.style.display = 'none';
                        });
    
                        const cardActions = document.createElement('div');
                        cardActions.classList.add('card-actions');
    
                        const viewButton = document.createElement('a');
                        viewButton.classList.add('card-action-button');
                        viewButton.textContent = 'View';
                        viewButton.href = `card-images/${decodedIdentifier}`;
                        viewButton.target = '_blank';
    
                        const deleteButton = document.createElement('button');
                        deleteButton.classList.add('card-action-button');
                        deleteButton.textContent = 'Delete';
                        deleteButton.addEventListener('click', () => deleteCard(decodedIdentifier));

                        const addToDeckButton = document.createElement('button');
                        addToDeckButton.classList.add('card-action-button');
                        addToDeckButton.textContent = 'Add to Deck';
                        addToDeckButton.addEventListener('click', () => addToDeck(decodedIdentifier));
                        cardActions.appendChild(addToDeckButton); 
    
    
                        cardActions.appendChild(viewButton);
                        cardActions.appendChild(deleteButton);
    
                        cardContainer.appendChild(cardImage);
                        cardContainer.appendChild(cardActions);
    
                        cardCanvas.appendChild(cardContainer);
                });
            })
            .catch(error => {
                console.error('Error fetching user cards:', error);
            });
    }
    
    
            // Call the function to load user's cards on page load
            loadUserCards();
            
            // Call loadUserCards when the page loads
            window.onload = loadUserCards;

            function deleteCard(fileName) {
                if (confirm('Are you sure you want to delete this card?')) {
                    fetch(`/cards/delete`, {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({ file_name: fileName }),
                    })
                    .then(response => response.json())
                    .then(data => {
                        if (data.success) {
                            alert('Card deleted successfully!');
                            location.reload(); // Reload the page to reflect the changes
                        } else {
                            alert('Error deleting card.');
                        }
                    })
                    .catch(error => {
                        console.error('Error deleting card:', error);
                    });
                }
            }

            // Populate the deck selector dropdown
            function populateDeckSelector(decks) {
                const deckSelector = document.getElementById('deckSelector');
            
                // Check if "None" option already exists
                const noneOption = deckSelector.querySelector('[value=""]');
                if (!noneOption) {
                    // Add the "None" option
                    const newNoneOption = document.createElement('option');
                    newNoneOption.value = '';
                    newNoneOption.textContent = 'None';
                    deckSelector.insertBefore(newNoneOption, deckSelector.firstChild);
                }
            
                // Add any missing deck options
                decks.forEach(deck => {
                    const option = deckSelector.querySelector(`[value="${deck}"]`);
                    if (!option) {
                        const newOption = document.createElement('option');
                        newOption.value = deck;
                        newOption.textContent = deck;
                        deckSelector.appendChild(newOption);
                    }
                });
            }
            

// Handle deck selection
function selectDeck() {
    const selectedDeck = document.getElementById('deckSelector').value;
    // Clear the deck parameter if "None" is selected
    const newUrl = selectedDeck === '' ? '/view-deck' : `/view-deck?deck=${encodeURIComponent(selectedDeck)}`;
    window.location.href = newUrl;
}
function addToDeck(fileName) {
    const deckName = prompt('Enter the deck name:');
    if (deckName) {
        fetch(`/cards/set-deck`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ file_name: fileName, deck: deckName }), // Use "deck" here
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Card added to deck successfully!');
                location.reload(); // Reload the page to reflect the changes
            } else {
                alert('Error adding card to deck.');
            }
        })
        .catch(error => {
            console.error('Error adding card to deck:', error);
        });
    }
}

// Fetch user's decks and populate the deck selector
function loadUserDecks() {
    fetch('/get-user-decks')
        .then(response => response.json())
        .then(decks => {
            populateDeckSelector(decks);
        })
        .catch(error => {
            console.error('Error fetching user decks:', error);
        });
}

// Call loadUserDecks when the page loads
window.onload = function() {
    loadUserDecks();
    populateDeckSelector([]);
};
package main

import (
	//"encoding/json"
	"fmt"
	"testing"
	"reflect"
	"sort"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
	
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var testCardEngine *xorm.Engine

func setupTestCardDatabase() error {
	// Connect to the test database
	var err error
	testCardEngine, err = xorm.NewEngine("mysql", "admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		return fmt.Errorf("failed to connect to the card database: %v", err)
	}

	// Sync the tables using XORM
	err = testCardEngine.Sync2(new(Card)) // Sync Card table
	if err != nil {
		return fmt.Errorf("failed to sync card tables: %v", err)
	}

	return nil
}

func TestSetDeck(t *testing.T) {
	fmt.Println("Testing SetDeck")

	// Initialize the database connection
	err := InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer CardEngine.Close()

	// Check if the test card exists in the database
	testFileName := "test_card.jpg"
	testCard, err := GetCardByFileName(testFileName)
	if err != nil {
		t.Fatalf("Failed to check test card in the database: %v", err)
	}

	// If the test card doesn't exist, create it with Deck "DeckA"
	if testCard == nil {
		_, err := NewCard("testuser", testFileName, "Test Card", "DeckA")
		if err != nil {
			t.Fatalf("Failed to create test card: %v", err)
		}
	}

	// Retrieve the test card by file name
	testCard, err = GetCardByFileName(testFileName)
	if err != nil {
		t.Fatalf("Failed to retrieve test card by file name: %v", err)
	}
	if testCard == nil {
		t.Fatalf("Test card not found")
	}

	// Record the current deck name
	originalDeck := testCard.Deck

	// Change the deck name based on the current value
	if testCard.Deck == "DeckA" {
		err = testCard.SetDeck("DeckB")
		if err != nil {
			t.Fatalf("Failed to set deck using SetDeck: %v", err)
		}
	} else {
		err = testCard.SetDeck("DeckA")
		if err != nil {
			t.Fatalf("Failed to set deck using SetDeck: %v", err)
		}
	}

	// Retrieve the updated test card
	updatedTestCard, err := GetCardByFileName(testFileName)
	if err != nil {
		t.Fatalf("Failed to retrieve updated test card by file name: %v", err)
	}
	if updatedTestCard == nil {
		t.Fatalf("Updated test card not found")
	}

	// Check if the deck name was properly updated
	expectedDeck := "DeckB" // Since we alternated the deck name
	if originalDeck == "DeckB" {
		expectedDeck = "DeckA" // Adjust the expected deck name if the original was "DeckB"
	}
	if updatedTestCard.Deck != expectedDeck {
		t.Errorf("Expected deck name '%s', but got '%s'", expectedDeck, updatedTestCard.Deck)
	}
}

func TestNewCard(t *testing.T) {
	fmt.Println("Testing NewCard")

	// Initialize the database connection
	err := InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer CardEngine.Close()

	// Delete the test card if it exists in the database
	_, err = CardEngine.Exec("DELETE FROM cards WHERE file_name = ?", "testcard1234.jpg")
	if err != nil {
		t.Fatalf("Failed to delete existing test card: %v", err)
	}

	// Create a new test card
	_, err = NewCard("testuser", "testcard1234.jpg", "Test Card", "DeckA")
	if err != nil {
		t.Fatalf("Failed to create test card: %v", err)
	}

	// Check if the test card exists in the database
	card := new(Card)
	has, err := CardEngine.Where("file_name = ?", "testcard1234.jpg").Get(card)
	if err != nil {
		t.Fatalf("Failed to retrieve test card from database: %v", err)
	}
	if !has {
		t.Errorf("Test card not found in the database")
	}

	// Delete the test card from the database
	_, err = CardEngine.Exec("DELETE FROM cards WHERE file_name = ?", "testcard1234.jpg")
	if err != nil {
		t.Fatalf("Failed to delete test card from database: %v", err)
	}
}


func TestGetCardsByUserID(t *testing.T) {
	fmt.Println("Testing GetCardsByUserID")

	// Initialize the database connection
	err := InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer CardEngine.Close()

	// Delete existing cards for test users
	_, err = CardEngine.Exec("DELETE FROM cards WHERE user_i_d IN (?, ?)", "testuser1", "testuser2")
	if err != nil {
		t.Fatalf("Failed to delete existing cards: %v", err)
	}

	// Create test cards for testuser1 and testuser2
	testUser1Cards := []string{"testcard1.jpg", "testcard2.jpg", "testcard3.jpg"}
	testUser2Cards := []string{"testcard4.jpg", "testcard5.jpg", "testcard6.jpg", "testcard7.jpg"}

	for _, cardFileName := range testUser1Cards {
		_, err := NewCard("testuser1", cardFileName, "Test Card", "DeckA")
		if err != nil {
			t.Fatalf("Failed to create test card for testuser1: %v", err)
		}
	}

	for _, cardFileName := range testUser2Cards {
		_, err := NewCard("testuser2", cardFileName, "Test Card", "DeckB")
		if err != nil {
			t.Fatalf("Failed to create test card for testuser2: %v", err)
		}
	}

	// Check cards for testuser1
	user1Cards, err := GetCardsByUserID("testuser1", "")
	if err != nil {
		t.Fatalf("Failed to retrieve cards for testuser1: %v", err)
	}
	if !reflect.DeepEqual(user1Cards, testUser1Cards) {
		t.Errorf("Retrieved cards for testuser1 do not match expected cards")
	}

	// Check cards for testuser2
	user2Cards, err := GetCardsByUserID("testuser2", "")
	if err != nil {
		t.Fatalf("Failed to retrieve cards for testuser2: %v", err)
	}
	if !reflect.DeepEqual(user2Cards, testUser2Cards) {
		t.Errorf("Retrieved cards for testuser2 do not match expected cards")
	}

	// Delete all cards for testuser1 and testuser2
	_, err = CardEngine.Exec("DELETE FROM cards WHERE user_i_d IN (?, ?)", "testuser1", "testuser2")
	if err != nil {
		t.Fatalf("Failed to delete test cards: %v", err)
	}
}


func TestGetCardByFileName(t *testing.T) {
	fmt.Println("Testing GetCardByFileName")

	// Initialize the database connection
	err := InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer CardEngine.Close()

	// Delete the test card if it exists in the database
	_, err = CardEngine.Exec("DELETE FROM cards WHERE file_name = ?", "testcard5678.jpg")
	if err != nil {
		t.Fatalf("Failed to delete existing test card: %v", err)
	}

	// Create a new test card
	_, err = NewCard("testuser", "testcard5678.jpg", "Test Card", "DeckA")
	if err != nil {
		t.Fatalf("Failed to create test card: %v", err)
	}

	// Get the test card by file name
	card, err := GetCardByFileName("testcard5678.jpg")
	if err != nil {
		t.Fatalf("Failed to retrieve test card: %v", err)
	}

	// Check if the card has the proper attributes
	if card.FileName != "testcard5678.jpg" || card.Name != "Test Card" || card.Deck != "DeckA" {
		t.Errorf("Test card attributes do not match expected values")
	}

	// Delete the test card from the database
	_, err = CardEngine.Exec("DELETE FROM cards WHERE file_name = ?", "testcard5678.jpg")
	if err != nil {
		t.Fatalf("Failed to delete test card from database: %v", err)
	}
}
func TestGetDecksByUserID(t *testing.T) {
	fmt.Println("Testing GetDecksByUserID")

	// Initialize the database connection
	err := InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer CardEngine.Close()

	// Delete the test cards with deck names 'TestDeck1', 'TestDeck2', 'TestDeck3'
	_, err = CardEngine.Exec("DELETE FROM cards WHERE deck IN (?, ?, ?)",
		"TestDeck1", "TestDeck2", "TestDeck3")
	if err != nil {
		t.Fatalf("Failed to delete existing test cards: %v", err)
	}

	// Create test cards
	testCards := []struct {
		UserID, DeckName string
	}{
		{"TestUser1", "TestDeck1"},
		{"TestUser1", "TestDeck1"},
		{"TestUser1", "TestDeck2"},
		{"TestUser1", "TestDeck2"},
		{"TestUser1", "TestDeck2"},
		{"TestUser2", "TestDeck1"},
		{"TestUser2", "TestDeck1"},
		{"TestUser2", "TestDeck1"},
		{"TestUser2", "TestDeck3"},
		{"TestUser2", "TestDeck3"},
	}

	for _, card := range testCards {
		_, err := NewCard(card.UserID, fmt.Sprintf("%s-%s.jpg", card.UserID, card.DeckName), "Test Card", card.DeckName)
		if err != nil {
			t.Fatalf("Failed to create test card: %v", err)
		}
		defer CardEngine.Exec("DELETE FROM cards WHERE file_name = ?", fmt.Sprintf("%s-%s.jpg", card.UserID, card.DeckName))
	}

	// Get unique decks for TestUser1
	decksUser1, err := GetDecksByUserID("TestUser1")
	if err != nil {
		t.Fatalf("Failed to retrieve decks for TestUser1: %v", err)
	}

	// Sort and compare the result with expected decks
	expectedDecksUser1 := []string{"TestDeck1", "TestDeck2"}
	sort.Strings(decksUser1)
	if !reflect.DeepEqual(decksUser1, expectedDecksUser1) {
		t.Errorf("Decks for TestUser1 do not match expected values")
	}

	// Get unique decks for TestUser2
	decksUser2, err := GetDecksByUserID("TestUser2")
	if err != nil {
		t.Fatalf("Failed to retrieve decks for TestUser2: %v", err)
	}

	// Sort and compare the result with expected decks
	expectedDecksUser2 := []string{"TestDeck1", "TestDeck3"}
	sort.Strings(decksUser2)
	if !reflect.DeepEqual(decksUser2, expectedDecksUser2) {
		t.Errorf("Decks for TestUser2 do not match expected values")
	}
}

func TestGetCardsByDeck(t *testing.T) {
    // Initialize the database connection
    err := InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }
    defer CardEngine.Close()

    // Clean up the existing test cards
    _, err = CardEngine.Exec("DELETE FROM cards WHERE user_i_d IN ('TestUser1', 'TestUser2')")
    if err != nil {
        t.Fatalf("Failed to clean up test cards: %v", err)
    }

    // Create test cards for TestUser1
    testCardsUser1 := []*Card{
        {UserID: "TestUser1", FileName: "testcard1.jpg", Name: "Card1", Deck: "TestDeck1"},
        {UserID: "TestUser1", FileName: "testcard2.jpg", Name: "Card2", Deck: "TestDeck1"},
    }

    for _, card := range testCardsUser1 {
        _, err := NewCard(card.UserID, card.FileName, card.Name, card.Deck)
        if err != nil {
            t.Fatalf("Failed to create test card: %v", err)
        }
    }

    // Create test cards for TestUser2
    testCardsUser2 := []*Card{
        {UserID: "TestUser2", FileName: "testcard3.jpg", Name: "Card3", Deck: "TestDeck1"},
        {UserID: "TestUser2", FileName: "testcard4.jpg", Name: "Card4", Deck: "TestDeck1"},
    }

    for _, card := range testCardsUser2 {
        _, err := NewCard(card.UserID, card.FileName, card.Name, card.Deck)
        if err != nil {
            t.Fatalf("Failed to create test card: %v", err)
        }
    }

    // Retrieve cards for TestUser1 and TestUser2
    user1Cards, err := GetCardsByDeck("TestUser1", "TestDeck1")
    if err != nil {
        t.Fatalf("Failed to retrieve cards for TestUser1: %v", err)
    }

    user2Cards, err := GetCardsByDeck("TestUser2", "TestDeck1")
    if err != nil {
        t.Fatalf("Failed to retrieve cards for TestUser2: %v", err)
    }

    // Verify the retrieved cards for TestUser1
    if len(user1Cards) != len(testCardsUser1) {
        t.Fatalf("Expected %d cards for TestUser1, but got %d", len(testCardsUser1), len(user1Cards))
    }

    for i, card := range user1Cards {
        if card.UserID != testCardsUser1[i].UserID || card.FileName != testCardsUser1[i].FileName ||
            card.Name != testCardsUser1[i].Name || card.Deck != testCardsUser1[i].Deck {
            t.Fatalf("Card mismatch for TestUser1")
        }
    }

    // Verify the retrieved cards for TestUser2
    if len(user2Cards) != len(testCardsUser2) {
        t.Fatalf("Expected %d cards for TestUser2, but got %d", len(testCardsUser2), len(user2Cards))
    }

    for i, card := range user2Cards {
        if card.UserID != testCardsUser2[i].UserID || card.FileName != testCardsUser2[i].FileName ||
            card.Name != testCardsUser2[i].Name || card.Deck != testCardsUser2[i].Deck {
            t.Fatalf("Card mismatch for TestUser2")
        }
    }
}


func teardownTestCardDatabase() error {
	// Clean up the test card database after the tests are run
	// For example, you can delete the test data or drop tables

	return nil
}

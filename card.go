
package main

import (
	"github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

// Card represents a user's uploaded card.
type Card struct {
    ID       int64  `xorm:"pk autoincr"`
    UserID   string `xorm:"not null"`
    FileName string `xorm:"not null"`
    Name     string
    Deck     string 
}

// TableName specifies the table name for the Card struct
func (Card) TableName() string {
    return "cards"
}


var CardEngine *xorm.Engine

func InitializeCardEngine(dbURL string) error {
	var err error
	CardEngine, err = xorm.NewEngine("mysql", dbURL)
	if err != nil {
		return err
	}

	// Sync the "Cards" table
	err = CardEngine.Sync2(new(Card))
	if err != nil {
		return err
	}

	return nil
}

// NewCard creates a new card record in the database.
// Tested: true
func NewCard(userID, fileName, name, deck string) (*Card, error) {
    card := &Card{
        UserID:   userID,
        FileName: fileName,
        Name:     name,
        Deck:     deck,
    }

    session := CardEngine.NewSession()
    defer session.Close()

    _, err := session.Insert(card)
    if err != nil {
        return nil, err
    }

    return card, nil
}
// GetCardsByUserID retrieves all card identifiers associated with a user.
// Tested: true
func GetCardsByUserID(userID string, deck string) ([]string, error) {
	//fmt.Println("Attempting to retrieve card identifiers for user " + userID)
    var cardIdentifiers []string
    query := CardEngine.Table("cards").Cols("file_name")

    if deck != ""  && deck != "None" {
        query = query.Where("user_i_d = ? AND deck = ?", userID, deck)
    } else {
        query = query.Where("user_i_d = ?", userID)
    }

    err := query.Find(&cardIdentifiers)
    if err != nil {
        fmt.Println("Failed to retrieve card identifiers ", err)
        return nil, err
    }

	// Decode the card identifiers
	decodedCardIdentifiers := make([]string, len(cardIdentifiers))
	for i, identifier := range cardIdentifiers {
		decodedIdentifier, err := url.QueryUnescape(identifier)
		if err != nil {
			fmt.Println("Error decoding identifier:", err)
			return nil, err
		}
		decodedCardIdentifiers[i] = decodedIdentifier
	}

	//fmt.Println("Retrieved card identifiers for user " + userID, decodedCardIdentifiers)
	return decodedCardIdentifiers, nil
}

// GetDecksByUserID retrieves a list of unique decks created by a user.
// Tested: true
func GetDecksByUserID(userID string) ([]string, error) {
    var decks []string
    err := CardEngine.Table("cards").Cols("deck").
        Where("user_i_d = ? AND deck != ?", userID, "None").
        GroupBy("deck").Find(&decks)
    if err != nil {
        return nil, err
    }
    return decks, nil
}


// DeleteCardByFileName deletes a card entry from the database based on the file name.
// Tested: false
func DeleteCardByFileName(fileName string) error {
    card := new(Card)
    _, err := CardEngine.Where("file_name = ?", fileName).Delete(card)
    if err != nil {
        return err
    }

	// Construct the path to the image file
	imagePath := filepath.Join("card-images", fileName)

	// Delete the image file
	err = os.Remove(imagePath)
	if err != nil {
		return err
	}

    return nil
}

// SetDeck updates the "deck" field for a specific card.
// Tested: true
func (c *Card) SetDeck(deckName string) error {
   // fmt.Println("SetDeck: "+deckName)
    session := CardEngine.NewSession()
    defer session.Close()

    _, err := session.Where("i_d = ?", c.ID).Cols("deck").Update(&Card{Deck: deckName})
    if err != nil {
        return err
    }

    return nil
}


// GetCardByFileName retrieves a single card by its file name.
// tested: true
func GetCardByFileName(fileName string) (*Card, error) {
    card := new(Card)
    has, err := CardEngine.Where("file_name = ?", fileName).Get(card)
    if err != nil {
        return nil, err
    }
    if !has {
        return nil, nil
    }
    return card, nil
}


// GetCardsByDeck retrieves all cards in the specified deck owned by the user.
func GetCardsByDeck(userID, deckName string) ([]*Card, error) {
    var cards []*Card
    err := CardEngine.Where("user_i_d = ? AND deck = ?", userID, deckName).Find(&cards)
    if err != nil {
        return nil, err
    }
    return cards, nil
}



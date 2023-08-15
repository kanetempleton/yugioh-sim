
package main

import (
	"github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/url"
)

// Card represents a user's uploaded card.
type Card struct {
    ID       int64  `xorm:"pk autoincr"`
    UserID   string `xorm:"not null"`
    FileName string `xorm:"not null"`
    Name     string
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
func NewCard(userID string, fileName, name string) (*Card, error) {
	card := &Card{
		UserID:   userID,
		FileName: fileName,
		Name:     name,
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
func GetCardsByUserID(userID string) ([]string, error) {
	fmt.Println("Attempting to retrieve card identifiers for user " + userID)
	var cardIdentifiers []string
	err := CardEngine.Table("cards").Cols("file_name").Where("user_i_d = ?", userID).Find(&cardIdentifiers)
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

	fmt.Println("Retrieved card identifiers for user " + userID, decodedCardIdentifiers)
	return decodedCardIdentifiers, nil
}



// Add more functions for managing cards, like deleting, updating, etc.

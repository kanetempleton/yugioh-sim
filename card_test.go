package main

import (
	//"encoding/json"
	"fmt"
	//"testing"

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

func teardownTestCardDatabase() error {
	// Clean up the test card database after the tests are run
	// For example, you can delete the test data or drop tables

	return nil
}

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var testEngine *xorm.Engine

func setupTestDatabase() error {
	// Connect to the test database
	var err error
	testEngine, err = xorm.NewEngine("mysql", "admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Sync the tables using XORM
	err = testEngine.Sync2(new(User)) // Sync User table
	if err != nil {
		return fmt.Errorf("failed to sync tables: %v", err)
	}

	return nil
}

func teardownTestDatabase() error {
	// Clean up the test database after the tests are run
	// For example, you can delete the test data or drop tables

	return nil
}

func TestCreateUser(t *testing.T) {
	fmt.Println("Testing CreateUser")
	// Set up the test database
	err := setupTestDatabase()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer teardownTestDatabase()

	// Create a new user
	user := &User{
		Username:   "testuser",
		Password:   "testpassword",
		Email:      "test@example.com",
		Privileges: 1,
	}

	// Check if the user already exists
	existingUser, err := GetUserByUsername(user.Username)
	if err != nil {
		t.Fatalf("Error querying database: %v", err)
	}

	// Delete the user if it already exists
	if existingUser != nil {
		_, err := Engine.Delete(existingUser)
		if err != nil {
			t.Fatalf("Failed to delete existing user: %v", err)
		}
	}


	err = CreateUser(user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Retrieve the created user from the test database
	testUser, err := GetUserByUsername("testuser")
	if err != nil {
		t.Fatalf("Failed to retrieve user from test database: %v", err)
	}

	// Compare the retrieved user with the original user
	if testUser.Username != user.Username ||
		testUser.Password != user.Password ||
		testUser.Email != user.Email ||
		testUser.Privileges != user.Privileges {
		t.Fatalf("Retrieved user does not match the original user")
	}

	// Delete the user at the end of the test
	_, err = Engine.Delete(user)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}
}

func TestMain(m *testing.M) {
	// Perform any setup or teardown tasks before/after all tests

	// Set up the test database
	err := setupTestDatabase()
	if err != nil {
		fmt.Printf("Failed to set up test database: %v", err)
		os.Exit(1)
	}

	// Initialize the Engine for test purposes
	err = Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		fmt.Printf("Failed to initialize test Engine: %v", err)
		os.Exit(1)
	}

	// Run the tests
	exitCode := m.Run()

	// Clean up the test database after all tests are run
	err = teardownTestDatabase()
	if err != nil {
		fmt.Printf("Failed to tear down test database: %v", err)
	}

	os.Exit(exitCode)
}


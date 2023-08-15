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

func TestDeleteUser(t *testing.T) {
	fmt.Println("Testing DeleteUser")
    // Initialize the database
    err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }
    defer Engine.Close()

    // Ensure the user 'testuser' is not in the database
    existingUser, err := GetUserByUsername("testuser")
    if err != nil {
        t.Fatalf("Failed to query database: %v", err)
    }
    if existingUser != nil {
        err := DeleteUser("testuser")
        if err != nil {
            t.Fatalf("Failed to delete user: %v", err)
        }
    }

    // Create a new user
    newUser := &User{
        Username: "testuser",
        Password: "testpassword",
        Email:    "testuser@example.com",
    }
    err = CreateUser(newUser)
    if err != nil {
        t.Fatalf("Failed to create user: %v", err)
    }

    // Delete the user
    err = DeleteUser("testuser")
    if err != nil {
        t.Fatalf("Failed to delete user: %v", err)
    }

    // Ensure the user is deleted from the database
    deletedUser, err := GetUserByUsername("testuser")
    if err != nil {
        t.Fatalf("Failed to query database: %v", err)
    }
    if deletedUser != nil {
        t.Fatalf("User was not deleted")
    }
}

func TestUpdateUser(t *testing.T) {
	fmt.Println("Testing UpdateUser")
    // Initialize the database connection
    err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }
    defer Engine.Close()

    // Ensure the test users don't exist before the test
    for _, username := range []string{"testuser1", "testuser2"} {
        _, err := GetUserByUsername(username)
        if err == nil {
            err = DeleteUser(username)
            if err != nil {
                t.Fatalf("Failed to delete test user: %v", err)
            }
        }
    }

    // Create test users
    testUsers := []*User{
        {Username: "testuser1", Password: "testpassword1", Email: "testuser1@example.com"},
        {Username: "testuser2", Password: "testpassword2", Email: "testuser2@example.com"},
    }

    for _, user := range testUsers {
        err := CreateUser(user)
        if err != nil {
            t.Fatalf("Failed to create test user: %v", err)
        }
    }

    // Update test users
    for _, user := range testUsers {
        user.Password = "newpassword"
        user.Email = "newemail@example.com"
        err := UpdateUser(user)
        if err != nil {
            t.Fatalf("Failed to update test user: %v", err)
        }
    }

    // Delete the test users
    for _, user := range testUsers {
        err := DeleteUser(user.Username)
        if err != nil {
            t.Fatalf("Failed to delete test user: %v", err)
        }
    }
}


func TestGetAllUsers(t *testing.T) {
	fmt.Println("Testing GetAllUsers")
    // Initialize the database connection
    err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }
    defer Engine.Close()

    // Ensure the test users don't exist before the test
    for _, username := range []string{"testuser1", "testuser2", "testuser3"} {
        _, err := GetUserByUsername(username)
        if err == nil {
            err = DeleteUser(username)
            if err != nil {
                t.Fatalf("Failed to delete test user: %v", err)
            }
        }
    }

    // Create test users
    testUsers := []*User{
        {Username: "testuser1", Password: "testpassword1", Email: "testuser1@example.com"},
        {Username: "testuser2", Password: "testpassword2", Email: "testuser2@example.com"},
        {Username: "testuser3", Password: "testpassword3", Email: "testuser3@example.com"},
    }

    for _, user := range testUsers {
        err := CreateUser(user)
        if err != nil {
            t.Fatalf("Failed to create test user: %v", err)
        }
    }

    // Retrieve all users
    users, err := GetAllUsers()
    if err != nil {
        t.Fatalf("Failed to retrieve all users: %v", err)
    }

    // Check if the number of retrieved users matches the number of test users
    if len(users) != len(testUsers) {
        t.Fatalf("Number of retrieved users does not match")
    }

    // Delete the test users
    for _, user := range testUsers {
        err := DeleteUser(user.Username)
        if err != nil {
            t.Fatalf("Failed to delete test user: %v", err)
        }
    }
}




func TestReadUserByEmail(t *testing.T) {
	fmt.Println("Testing ReadUserByEmail")
    // Initialize the database connection
    err := Initialize("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }
    defer Engine.Close()

    // Ensure the test user doesn't exist before the test
    _, err = GetUserByEmail("testuser@example.com")
    if err == nil {
        err = DeleteUser("testuser")
        if err != nil {
            t.Fatalf("Failed to delete test user: %v", err)
        }
    }

    // Create a test user
    testUser := &User{
        Username: "testuser",
        Password: "testpassword",
        Email:    "testuser@example.com",
    }

    err = CreateUser(testUser)
    if err != nil {
        t.Fatalf("Failed to create test user: %v", err)
    }

    // Retrieve the user by email
    retrievedUser, err := GetUserByEmail("testuser@example.com")
    if err != nil {
        t.Fatalf("Failed to retrieve user by email: %v", err)
    }

    // Check if the retrieved user matches the test user
    if retrievedUser.Username != testUser.Username ||
        retrievedUser.Password != testUser.Password ||
        retrievedUser.Email != testUser.Email ||
        retrievedUser.Privileges != testUser.Privileges {
        t.Fatalf("Retrieved user does not match test user")
    }

    // Delete the test user
    err = DeleteUser("testuser")
    if err != nil {
        t.Fatalf("Failed to delete test user: %v", err)
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

	err = setupTestCardDatabase()
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

	// Initialize the CardEngine for test purposes
	err = InitializeCardEngine("admin:obviouspassword@tcp(localhost:3306)/yugiohgo_test")
	if err != nil {
		fmt.Printf("Failed to initialize test CardEngine: %v", err)
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



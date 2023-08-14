package main

import (
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)

// User represents a user account.
type User struct {
    ID        int64  `xorm:"pk autoincr"`
    Username  string `xorm:"unique"`
    Password  string
    Email     string
    Privileges int
}

// Engine is the global XORM engine.
var Engine *xorm.Engine

// Initialize initializes the XORM engine and syncs the "Users" table.
func Initialize(dbURL string) error {
    var err error
    Engine, err = xorm.NewEngine("mysql", dbURL)
    if err != nil {
        return err
    }

    // Sync the "Users" table
    err = Engine.Sync2(new(User))
    if err != nil {
        return err
    }

    return nil
}


// CreateUser inserts a new user into the database.
func CreateUser(user *User) error {
    _, err := Engine.Insert(user)
    return err
}

// GetUserByUsername retrieves a user by username from the database.
func GetUserByUsername(username string) (*User, error) {
    user := new(User)
    has, err := Engine.Where("username = ?", username).Get(user)
    if err != nil {
        return nil, err
    }
    if !has {
        return nil, nil
    }
    return user, nil
}

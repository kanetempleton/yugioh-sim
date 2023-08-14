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

// DeleteUser deletes a user from the database.
func DeleteUser(username string) error {
    _, err := Engine.Where("username = ?", username).Delete(&User{})
    return err
}

func UpdateUser(user *User) error {
    _, err := Engine.Where("username = ?", user.Username).Update(user)
    return err
}

func GetAllUsers() ([]User, error) {
    var users []User
    err := Engine.Find(&users)
    if err != nil {
        return nil, err
    }
    return users, nil
}

func GetUserByEmail(email string) (*User, error) {
    user := new(User)
    has, err := Engine.Where("email = ?", email).Get(user)
    if err != nil {
        return nil, err
    }
    if !has {
        return nil, nil
    }
    return user, nil
}


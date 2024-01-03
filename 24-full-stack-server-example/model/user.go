package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
    gorm.Model
    Name  string
    Email string
}


func ValidateUser(user *User) error {
    if len(strings.TrimSpace(user.Name)) < 3 {
        return errors.New("user's name is too short, must be at least 3 characters")
    }
    if !strings.Contains(user.Email, "@") { // This is a very basic check, consider using a better validation for email
        return errors.New("invalid email format")
    }
    return nil
}

var ErrUserNotFound = errors.New("user not found")

// GetUserById fetches a user by their ID from the database
func GetUserById(db *gorm.DB, id int) (User, error) {
    var user User
    result := db.First(&user, id)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return User{}, ErrUserNotFound
    }
    return user, result.Error
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user User) error {
    result := db.Create(&user)
    return result.Error
}

// GetAllUsers retrieves users from the database with pagination and returns total count
func GetAllUsers(db *gorm.DB, page, itemsPerPage int) ([]User, int64, error) {
    var users []User
    var totalCount int64

    db.Model(&User{}).Count(&totalCount) // Get total count of users

    offset := (page - 1) * itemsPerPage
    result := db.Offset(offset).Limit(itemsPerPage).Find(&users)
    return users, totalCount, result.Error
}


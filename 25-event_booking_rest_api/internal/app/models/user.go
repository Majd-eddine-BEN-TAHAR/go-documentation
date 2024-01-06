package models

// User represents the user model in our application.
// It contains the username, password, and email of the user.
type User struct {
	ID       int  
	Username string 
	Password string
	Email	 string
}
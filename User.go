package main

// User object
type User struct {
	UserID   string `json:"user_id" description:"Unique user identifier"`
	Username string `json:"username" description:"Username of user"`
}

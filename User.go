package main

type User struct {
	UserId   string `json:"user_id" description:"Unique user identifier"`
	Username string `json:"username" description:"Username of user"`
}

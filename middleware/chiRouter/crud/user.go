package main

import "errors"

// User data model
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// User fixture data
var users = []*User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

func dbGetUser(id int64) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}

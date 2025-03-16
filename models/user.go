package models

import (
	"arena"
	"github.com/emreisler/go-arena-tracking/tracking"
)

// User struct (private)
type User struct {
	ID   int
	Name string
	Tags []string
}

func NewUser(ar *arena.Arena, id int, name string, tags []string) *User {
	u := arena.New[User](ar)
	u.ID = id
	u.Name = name
	u.Tags = tags

	tracking.TrackHeapAlloc("User", u)
	return u
}

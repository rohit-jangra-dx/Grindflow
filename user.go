package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID		uint32
	CreatedAt	string
	Name		string
	Projects	[]Project
	Tags		[]Tag
}

// creates a new user
func NewUser(name string) *User {
	
	return &User{
		UUID: uuid.New().ID(),
		Name: name,
		CreatedAt: time.Now().Format(time.RFC3339),
		Projects: []Project{},
		Tags: []Tag{},
	}
	
}
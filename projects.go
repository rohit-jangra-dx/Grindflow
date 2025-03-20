// contains code related to projects and tasks
package main

import (
	"time"

	"github.com/google/uuid"
)




type Project struct {
	Name 		string
	UUID		uint32
	Tasks		[]Task
	CreatedAt	string	
}

type Task struct {
	UUID		uint32
	Title 		string
	Description	string
	CreatedAt	string	
	Owner		uint32
	ScheduledAt	string	
	Tags		[]Tag 
}

type Tag struct {
	UUID		uint32
	Name		string
}

/*  project related methods */
// Create A New Project
func NewProject(name string) Project {
	return Project {
		UUID: uuid.New().ID(),
		Name: name,
		Tasks: []Task{},
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

/*  task related methods */

func NewTask(title string, desc string, owner uint32, scheduledAt string, tags []Tag ) Task {
	return Task{
		UUID: uuid.New().ID(),
		Title: title,
		Description: desc,
		Owner: owner,
		ScheduledAt: scheduledAt,
		Tags: tags,
	}
}

/* tags related stuff */
func NewTag(name string) Tag {
	return Tag {
		UUID: uuid.New().ID(),
		Name: name,
	}
}
package main

import (
	"context"
	"fmt"
	"slices"
)

// App struct
type App struct {
	ctx context.Context
	User *User
}

// NewApp creates a new App application struct 
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// CreateUser is a wrapper for NewUser helper method
func (a *App) CreateUser(name string) (string, error) {
	
	if a.User != nil {
		return "", fmt.Errorf("User already exists: %s", a.User.Name)
	}
	
	a.User = NewUser(name)
	return fmt.Sprintf("User %s created successfully!", a.User.Name), nil

}

func (a *App) GetUser() (User, error) {
	if a.User != nil {
		return *a.User, nil
	}
	
	return User{}, fmt.Errorf("User hasn't being created yet")
}

func (a *App) CreateNewProject(name string) (string, error) {
	
	// check if the project with same name exists or not
	projectAlreadyExists := slices.ContainsFunc(a.User.Projects, func(p Project) bool {
		return p.Name == name
	})
	
	if projectAlreadyExists {
		return "", fmt.Errorf("project already exists: %s", name)
	}
	
	a.User.Projects = append(a.User.Projects, NewProject(name))
	return fmt.Sprintf("Project %s created successfully!", name), nil
}
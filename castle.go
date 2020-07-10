package main

// Application ...
type Application struct {
  name string
  Roles map[string]*Role
  Profiles map[string]*Profile
}

// Role ...
type Role struct {}

// Profile ...
type Profile struct {}

// NewApplication ...
func NewApplication(name string) (*Application, error) {
  // TODO: check name matches a-zA-Z

  return &Application{name}
}

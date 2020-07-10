package castle

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

var applications map[string]*Application

func init() {
	applications = make(map[string]*Application)
}

type ApplicationInterface interface {
	NewPermission() *Permission
	NewRole(string, ...*Permission) RoleInterface
	GetRole(string) (RoleInterface, error)
	String() string
}

// Application ...
type Application struct {
	name        string
	permissions map[string]*Permission
	roles       map[string]*Role
}

// NewPermission ...
func (a *Application) NewPermission() *Permission {
	permissionUUID := uuid.New()
	name := permissionUUID.String()

	a.permissions[name] = &Permission{
		UUID: permissionUUID,
	}

	return a.permissions[name]
}

// NewRole ...
func (a *Application) NewRole(name string, permissions ...*Permission) RoleInterface {
	a.roles[name] = &Role{
		Name:        fmt.Sprintf("%s.%s", a.name, name),
		Permissions: permissions,
	}

	return a.roles[name]
}

// GetRole ...
func (a *Application) GetRole(name string) (RoleInterface, error) {
	for key, value := range a.roles {
		if key == name {
			return value, nil
		}
	}

	return nil, fmt.Errorf("cannot retrieve role: %s", name)
}

func (a *Application) String() string {
	return a.name
}

// NewApplication ...
func NewApplication(name string) (ApplicationInterface, error) {
	// TODO: check name matches a-zA-Z

	applications[name] = &Application{
		name:        name,
		permissions: map[string]*Permission{},
		roles:       map[string]*Role{},
	}

	return applications[name], nil
}

// GetRole ...
func GetRole(role string) (RoleInterface, error) {
	var app *Application = nil

	segment := strings.Split(role, ".")

	for name, application := range applications {
		if name == segment[0] {
			app = application
			break
		}
	}

	if app == nil {
		return nil, fmt.Errorf("cannot retrieve application: %s", segment[0])
	}

	for name, role := range app.roles {
		if name == segment[1] {
			return role, nil
		}
	}

	return nil, fmt.Errorf("cannot retrieve role: %s", segment[1])
}

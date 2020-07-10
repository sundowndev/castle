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
func (c *Application) NewPermission() *Permission {
	permissionUUID := uuid.New()
	name := permissionUUID.String()

	c.permissions[name] = &Permission{
		UUID: permissionUUID,
	}

	return c.permissions[name]
}

// NewRole ...
func (c *Application) NewRole(name string, permissions ...*Permission) RoleInterface {
	c.roles[name] = &Role{
		Name:        fmt.Sprintf("%s.%s", c.name, name),
		Permissions: permissions,
	}

	return c.roles[name]
}

// GetRole ...
func (c *Application) GetRole(name string) (RoleInterface, error) {
	for key, value := range c.roles {
		if key == name {
			return value, nil
		}
	}

	return nil, fmt.Errorf("cannot retrieve role: %s", name)
}

func (c *Application) String() string {
	return c.name
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

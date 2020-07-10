package castle

import (
	"testing"

	assertion "github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should return valid application name", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		app, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		assert.Equal("myapp", app.String(), "they should be equal")
	})

	t.Run("should create a role with a valid name", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		app, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		permission1 := app.NewPermission()

		app.NewRole("role", permission1)

		role, err := GetRole("myapp.role")
		if err != nil {
			panic(err)
		}

		assert.Equal("myapp.role", role.String(), "they should be equal")
	})

	t.Run("should have only 1 permission", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		app, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		permission1 := app.NewPermission()
		permission2 := app.NewPermission()

		app.NewRole("role", permission1)

		role, err := GetRole("myapp.role")
		if err != nil {
			panic(err)
		}

		assert.Equal(true, role.HasPermission(permission1), "they should be equal")
		assert.Equal(false, role.HasPermission(permission2), "they should be equal")
	})

	t.Run("should get role from container", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		app, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		app.NewRole("role")

		role, err := app.GetRole("role")
		if err != nil {
			panic(err)
		}

		assert.Equal("myapp.role", role.String(), "they should be equal")
	})

	t.Run("should fail to retrieve application", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		_, err := GetRole("myapp.test")

		assert.Equal("cannot retrieve application: myapp", err.Error(), "they should be equal")
	})

	t.Run("should fail to retrieve role globally", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		_, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		_, err = GetRole("myapp.test")

		assert.Equal("cannot retrieve role: test", err.Error(), "they should be equal")
	})

	t.Run("should fail to retrieve role from container", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		app, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		_, err = app.GetRole("test")

		assert.Equal("cannot retrieve role: test", err.Error(), "they should be equal")
	})

	t.Run("should inherit from another role", func(t *testing.T) {
		applications = make(map[string]*Application) // Reset applications

		app, err := NewApplication("myapp")
		if err != nil {
			panic(err)
		}

		permission1 := app.NewPermission()
		permission2 := app.NewPermission()

		role1 := app.NewRole("role", permission1)
		role2 := app.NewRole("role", permission2)

		assert.Equal(false, role2.HasPermission(permission1), "they should be equal")

		role2.InheritFrom(role1)

		assert.Equal(true, role2.HasPermission(permission1), "they should be equal")
	})
}

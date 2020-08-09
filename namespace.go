package castle

import "fmt"

// Namespace refers to a resource of an application.
type Namespace struct {
	name string
	app  *Application
}

// NewNamespace creates a namespace
func (ns *Namespace) NewNamespace(name string) *Namespace {
	fullName := fmt.Sprintf("%s.%s", ns.name, name)

	ns.app.namespaces[fullName] = &Namespace{
		name: fullName,
		app:  ns.app,
	}

	return ns.app.namespaces[fullName]
}

// NewScope creates a new scope using the given name
// This function overrides any duplicated usage
func (ns *Namespace) NewScope(name string) *Scope {
	fullName := fmt.Sprintf("%s.%s", ns.name, name)

	ns.app.scopes[name] = &Scope{
		namespace: ns,
		name:      fullName,
	}

	return ns.app.scopes[name]
}

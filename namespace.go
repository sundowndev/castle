package castle

// Namespace refers to a resource of an application.
type Namespace struct {
	name string
	app  *Application
}

// NewNamespace creates a namespace
func (ns *Namespace) NewNamespace(name string) *Namespace {
	ns.app.namespaces[name] = &Namespace{
		name: name,
		app:  ns.app,
	}

	return ns.app.namespaces[name]
}

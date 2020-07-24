package castle

// Namespace ...
type Namespace struct {
	name string
	app  *Application
}

// Score ...
type Scope struct {
	namespace *Namespace
	name      string
}

func (s *Scope) String() string {
	return s.name
}

func (n *Namespace) NewScope(name string) *Scope {
	n.app.scopes[name] = &Scope{
		namespace: n,
		name:      name,
	}

	return n.app.scopes[name]
}

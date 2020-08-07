package castle

// Scope defines a resource associated to a namespace and granted to tokens
type Scope struct {
	namespace *Namespace
	name      string
}

// String returns a string representation of the scope
func (s *Scope) String() string {
	return s.name
}

// NewScope creates a new scope using the given name
// This function overrides any duplicated usage
func (n *Namespace) NewScope(name string) *Scope {
	n.app.scopes[name] = &Scope{
		namespace: n,
		name:      name,
	}

	return n.app.scopes[name]
}

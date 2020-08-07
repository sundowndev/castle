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

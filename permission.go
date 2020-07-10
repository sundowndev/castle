package castle

import "github.com/google/uuid"

type callback func (...interface{}) bool

// Permission ...
type Permission struct {
	UUID     uuid.UUID
	Callback callback
}

func (p *Permission) String() string {
	return p.UUID.String()
}
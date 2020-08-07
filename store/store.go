package store

import (
	"sync"
	"time"
)

type Store interface {
	GetKey(string) (string, error)
	SetKey(string, string, time.Time) error
	RemoveKey(string) (bool, error)
	Flush() error
}

// NewLocalStore creates a new local store to be used in an application
func NewLocalStore() *LocalStore {
	return &LocalStore{
		store: map[string]string{},
		m:     &sync.RWMutex{},
	}
}

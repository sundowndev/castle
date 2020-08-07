package store

import (
	"fmt"
	"sync"
	"time"
)

// LocalStore is an in-memory key-value storage for testing
type LocalStore struct {
	m     *sync.RWMutex
	store map[string]string
}

// GetKey reads a key in the store
func (s *LocalStore) GetKey(key string) (string, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	if v, ok := s.store[key]; ok {
		return v, nil
	}

	return "", fmt.Errorf("key not found: %s", key)
}

// SetKey writes a key in the store
// Warning: writing on a key that already exists will cause an error.
// Delete it first.
func (s *LocalStore) SetKey(key string, value string, expiration time.Time) error {
	s.m.Lock()
	defer s.m.Unlock()

	// Ensure key doesn't exists yet
	if _, ok := s.store[key]; ok {
		return fmt.Errorf("key already exist: %s", key)
	}

	s.store[key] = value

	time.AfterFunc(expiration.Sub(time.Now()), func() {
		_, _ = s.RemoveKey(key)
	})

	return nil

}

// RemoveKey removes a key in the store
func (s *LocalStore) RemoveKey(key string) (bool, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var removed bool

	if _, ok := s.store[key]; ok {
		removed = true
	}

	delete(s.store, key)

	return removed, nil
}

// Flush removes every single keys in the store
func (s *LocalStore) Flush() error {
	s.m.Lock()
	defer s.m.Unlock()

	for k := range s.store {
		delete(s.store, k)
	}

	return nil
}

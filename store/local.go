package store

import (
	"fmt"
	"sync"
	"time"
)

type LocalStore struct {
	m     *sync.RWMutex
	store map[string]string
}

func (s *LocalStore) GetKey(key string) (string, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	if v, ok := s.store[key]; ok {
		return v, nil
	}

	return "", fmt.Errorf("key not found: %s", key)
}

func (s *LocalStore) SetKey(key string, value string, expiration time.Time) error {
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

func (s *LocalStore) RemoveKey(key string) (bool, error) {
	var removed bool

	if _, ok := s.store[key]; ok {
		removed = true
	}

	delete(s.store, key)

	return removed, nil
}

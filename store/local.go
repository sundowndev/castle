package store

import "fmt"

type Store interface {
	GetKey(string) (string,error)
	SetKey(string, string, int) error
	RemoveKey(string) (bool, error)
}

type LocalStore struct {
	Store map[string]string
}

func (s *LocalStore) GetKey(key string) (string, error) {
	if v, ok := s.Store[key]; ok {
		return v, nil
	}

	return "", fmt.Errorf("key not found: %s", key)
}

func (s *LocalStore) SetKey(key string, value string, expiration int) error {
	// Ensure key doesn't exists yet
	if _, ok := s.Store[key]; !ok {
		s.Store[key] = value

		return nil
	}

	return fmt.Errorf("key already exist: %s. Remove it first", key)
}

func (s *LocalStore) RemoveKey(key string) (bool, error) {
	var removed bool

	if _, ok := s.Store[key]; !ok {
		removed = true
	}

	delete(s.Store, key)

	return removed, nil
}

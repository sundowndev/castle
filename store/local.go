package store

import (
	"fmt"
	"time"
)

type Store interface {
	GetKey(string) (string, error)
	SetKey(string, string, time.Time) error
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

func (s *LocalStore) SetKey(key string, value string, expiration time.Time) error {
	// Ensure key doesn't exists yet
	if _, ok := s.Store[key]; ok {
		return fmt.Errorf("key already exist: %s", key)
	}

	s.Store[key] = value

	go s.removeWhenExpired(key, expiration)

	return nil

}

func (s *LocalStore) RemoveKey(key string) (bool, error) {
	var removed bool

	if _, ok := s.Store[key]; ok {
		removed = true
	}

	delete(s.Store, key)

	return removed, nil
}

func (s *LocalStore) removeWhenExpired(key string, expiration time.Time) error {
	tick := time.Tick(1 * time.Second)

	for {
		select {
		case <-tick:
			if expiration.Before(time.Now()) {
				_, err := s.RemoveKey(key)
				if err != nil {
					return err
				}

				return nil
			}
		}
	}
}

package store

import (
	"fmt"
	"sync"
	"time"
)

type LocalStore struct {
	Store *sync.Map
}

func (s *LocalStore) GetKey(key string) (string, error) {
	if v, ok := s.Store.Load(key); ok {
		return v.(string), nil
	}

	return "", fmt.Errorf("key not found: %s", key)
}

func (s *LocalStore) SetKey(key string, value string, expiration time.Time) error {
	// Ensure key doesn't exists yet
	if _, ok := s.Store.Load(key); ok {
		return fmt.Errorf("key already exist: %s", key)
	}

	s.Store.Store(key, value)

	go s.removeWhenExpired(key, expiration)

	return nil

}

func (s *LocalStore) RemoveKey(key string) (bool, error) {
	var removed bool

	if _, ok := s.Store.Load(key); ok {
		removed = true
	}

	s.Store.Delete(key)

	return removed, nil
}

func (s *LocalStore) removeWhenExpired(key string, expiration time.Time) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
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

package store

import (
	"github.com/sundowndev/castle"
	"testing"
	"time"

	assertion "github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should set a key", func(t *testing.T) {
		store := castle.NewLocalStore()

		err := store.SetKey("hello", "word", time.Now().Add(1*time.Minute))
		assert.Nil(nil, err)

		v, err := store.GetKey("hello")
		assert.Nil(nil, err)

		assert.Equal("word", v)
	})

	t.Run("should set a key that already exists", func(t *testing.T) {
		store := castle.NewLocalStore()

		_ = store.SetKey("hello", "word", time.Now().Add(1*time.Minute))

		err := store.SetKey("hello", "word", time.Now().Add(1*time.Minute))

		assert.EqualError(err, "key already exist: hello")
	})

	t.Run("should set then remove a key", func(t *testing.T) {
		store := castle.NewLocalStore()

		err := store.SetKey("hello", "word", time.Now().Add(1*time.Minute))
		assert.Nil(nil, err)

		removed, err := store.RemoveKey("hello")
		assert.Nil(nil, err)

		assert.Equal(true, removed)
	})

	t.Run("should remove non-existent key", func(t *testing.T) {
		store := castle.NewLocalStore()

		removed, err := store.RemoveKey("hello")
		assert.Nil(nil, err)

		assert.Equal(false, removed)
	})

	t.Run("should get an expired key", func(t *testing.T) {
		store := castle.NewLocalStore()

		err := store.SetKey("hello", "word", time.Now())
		assert.Nil(nil, err)

		time.Sleep(2 * time.Second)

		v, err := store.GetKey("hello")
		assert.Equal("", v)
		assert.EqualError(err, "key not found: hello")
	})
}

package castle_test

import (
	"github.com/google/uuid"
	"github.com/sundowndev/castle"
	"github.com/sundowndev/castle/store"
	"testing"
	"time"

	assertion "github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should create a valid token", func(t *testing.T) {
		app := castle.NewApp(store.NewLocalStore())

		ns := app.NewNamespace("repositories")

		read := ns.NewScope("read_repository")

		token, err := app.NewToken("myrepo", time.Now().Add(1*time.Minute), read)
		assert.Nil(err)

		assert.Equal("myrepo", token.Name, "they should be equal")

		_, err = uuid.Parse(token.String())
		assert.Nil(err)
	})

	t.Run("should retrieve a token", func(t *testing.T) {
		app := castle.NewApp(store.NewLocalStore())

		ns := app.NewNamespace("repositories")

		read := ns.NewScope("read_repository")

		token, err := app.NewToken("myrepo", time.Now().Add(1*time.Minute), read)
		assert.Nil(err)

		token2, err := app.GetToken(token.String())
		assert.Nil(err)

		assert.Equal(token.Name, token2.Name)
		assert.Equal(read.String(), token.Scopes[0])
	})

	t.Run("should revoke a token", func(t *testing.T) {
		app := castle.NewApp(store.NewLocalStore())

		ns := app.NewNamespace("repositories")

		read := ns.NewScope("read_repository")

		token, err := app.NewToken("myrepo", time.Now().Add(1*time.Minute), read)
		assert.Nil(err)

		err = app.RevokeToken(token.String())
		assert.Nil(err)

		_, err = app.GetToken(token.String())
		assert.Errorf(err, "key not found: %s", token.String())
	})

	t.Run("should set rate limit on token", func(t *testing.T) {
		app := castle.NewApp(store.NewLocalStore())

		ns := app.NewNamespace("repositories")

		read := ns.NewScope("read_repository")

		token, err := app.NewToken("myrepo", time.Now().Add(1*time.Minute), read)
		assert.Nil(err)

		err = app.RateLimitFunc(token, func(rate int) int {
			return 100
		})
		assert.Nil(err)

		rate, err := app.GetRateLimit(token)
		assert.Nil(err)

		assert.Equal(100, rate)
	})

	t.Run("should set rate limit on token but not going lower than 0", func(t *testing.T) {
		app := castle.NewApp(store.NewLocalStore())

		ns := app.NewNamespace("repositories")

		read := ns.NewScope("read_repository")

		token, err := app.NewToken("myrepo", time.Now().Add(1*time.Minute), read)
		assert.Nil(err)

		err = app.RateLimitFunc(token, func(rate int) int {
			return rate - 10
		})
		assert.Nil(err)

		rate, err := app.GetRateLimit(token)
		assert.Nil(err)

		assert.Equal(0, rate)
	})
}

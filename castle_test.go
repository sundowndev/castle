package castle_test

import (
	"github.com/sundowndev/castle"
	"github.com/sundowndev/castle/store"
	"regexp"
	"testing"

	assertion "github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should create a valid token", func(t *testing.T) {
		app := castle.NewApp(&store.LocalStore{
			Store: make(map[string]string),
		})

		ns := app.NewNamespace("repositories")

		read := ns.NewScope("read_repository")

		token, err := app.NewToken("myrepo", 3600, read)
		assert.Nil(err)

		assert.Equal( "myrepo",token.Name, "they should be equal")
		assert.Regexp( regexp.MustCompile("([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}){1}"), token.String(),"they should be equal")
	})
}

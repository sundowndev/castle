package castle

import (
	"github.com/google/uuid"
	assertion "github.com/stretchr/testify/assert"
	"testing"
)

func TestToken(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should serialize token properly", func(t *testing.T) {
		token := &Token{
			uuid:   uuid.New(),
			Name:   "mytoken",
			Scopes: []string{"read_repository"},
			RateLimit: -1,
		}

		json, err := token.Serialize()

		assert.Nil(err)
		assert.Equal("{\"name\":\"mytoken\",\"scopes\":[\"read_repository\"],\"RateLimit\":-1}", json)
	})

	t.Run("should check scope on token", func(t *testing.T) {
		read := &Scope{name: "read", namespace: nil}
		write := &Scope{name: "write", namespace: nil}

		token := &Token{
			uuid:   uuid.New(),
			Name:   "mytoken",
			Scopes: []string{"read"},
		}

		assert.True(token.HasScope(read))
		assert.False(token.HasScope(write))
	})
}

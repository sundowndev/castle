package castle

import (
	"github.com/google/uuid"
	assertion "github.com/stretchr/testify/assert"
	"testing"
)

func TestToken(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should serialize token properly", func(t *testing.T) {
		scope := "read"

		token := &Token{
			uuid:      uuid.New(),
			Name:      "mytoken",
			Scopes:    []string{scope},
			RateLimit: -1,
		}

		json, err := token.Serialize()

		assert.Nil(err)
		assert.Equal("{\"name\":\"mytoken\",\"scopes\":[\""+scope+"\"]}", json)
	})

	t.Run("should deserialize token properly", func(t *testing.T) {
		token, err := deserialize("0785ed1c-b5e9-4a1b-b972-45065d8ad660", []byte("{\"name\":\"mytoken\",\"scopes\":[\"read\"]}"))
		assert.Nil(err)

		assert.Equal("0785ed1c-b5e9-4a1b-b972-45065d8ad660", token.String())
		assert.Equal("mytoken", token.Name)
		assert.Equal([]string{"read"}, token.Scopes)
		assert.Equal(Rate(0), token.RateLimit)
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

package castle

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Rate int

type Token struct {
	uuid      uuid.UUID
	Name      string   `json:"name"`
	Scopes    []string `json:"scopes"`
	RateLimit Rate     `json:"-" default:"-1"`
	expiresAt time.Time
}

// String returns the token's value as a string.
// That value is always a valid UUID.
func (t *Token) String() string {
	return t.uuid.String()
}

// Serialize serializes the token to a JSON string
func (t *Token) Serialize() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// HasScope checks if the current token has access to the given scope
func (t *Token) HasScope(scope *Scope) bool {
	for _, v := range t.Scopes {
		if v == scope.String() {
			return true
		}
	}
	return false
}

func deserialize(key string, j []byte) (*Token, error) {
	u, err := uuid.Parse(key)
	if err != nil {
		return nil, err
	}

	var t = Token{
		uuid: u,
	}

	err = json.Unmarshal(j, &t)
	if err != nil {
		return &t, err
	}

	return &t, nil
}

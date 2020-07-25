package castle

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Token struct {
	uuid      uuid.UUID `json:"-"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Scopes    []string  `json:"scopes"`
	RateLimit int       `json:"rate_limit" default:"-1"`
	expiresAt time.Time `json:"-"`
}

func (t *Token) String() string {
	return t.uuid.String()
}

func (t *Token) SetRateLimit(rate int) {
	t.RateLimit = rate
}

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

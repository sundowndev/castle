package castle

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Token struct {
	uuid      uuid.UUID `json:"-"`
	Name      string    `json:"name"`
	Namespace string `json:"namespace"`
	Scopes    []string  `json:"scopes"`
	Expiration int `json:"expiration"`
}

func (t *Token) String() string {
	return t.uuid.String()
}

func (t *Token) Serialize() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (t *Token) HasScope(scope *Scope) bool {
	return false
}
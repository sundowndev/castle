package castle

import (
	"github.com/google/uuid"
	"github.com/sundowndev/castle/store"
	"strconv"
	"time"
)

// LocalStore is an in-memory key-value storage for testing
type LocalStore = store.LocalStore

// Application defines an entry point for your web application
type Application struct {
	store      store.Store
	namespaces map[string]*Namespace
	scopes     map[string]*Scope
}

const defaultRateLimit int = -1
const tokenKeySuffix string = ":token"
const rateLimitKeySuffix string = ":rate"

// NewApp creates an application object with a key/value storage, scopes and namespaces
func NewApp(s store.Store) *Application {
	return &Application{
		store:      s,
		namespaces: make(map[string]*Namespace),
		scopes:     make(map[string]*Scope),
	}
}

// NewNamespace creates a namespace
func (a *Application) NewNamespace(name string) *Namespace {
	a.namespaces[name] = &Namespace{
		name: name,
		app:  a,
	}

	return a.namespaces[name]
}

// NewToken creates a new UUID token
func (a *Application) NewToken(name string, expiration time.Time, scopes ...*Scope) (*Token, error) {
	var scopesAsString []string

	for _, v := range scopes {
		scopesAsString = append(scopesAsString, v.String())
	}

	t := &Token{
		uuid:   uuid.New(),
		Name:   name,
		Scopes: scopesAsString,
	}

	err := a.RateLimitFunc(t, func(_ int) int {
		return defaultRateLimit
	})
	if err != nil {
		return nil, err
	}

	json, err := t.Serialize()
	if err != nil {
		return nil, err
	}

	err = a.store.SetKey(t.String()+tokenKeySuffix, json, expiration)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetToken retrieve a token from its value
func (a *Application) GetToken(token string) (*Token, error) {
	json, err := a.store.GetKey(serializeTokenKey(token))
	if err != nil {
		return nil, err
	}

	t, err := deserialize(token, []byte(json))
	if err != nil {
		return nil, err
	}

	return t, nil
}

// RevokeToken permanently delete a token
func (a *Application) RevokeToken(token string) error {
	_, err := a.store.RemoveKey(serializeTokenKey(token))

	return err
}

// RateLimitFunc mutate the rate limit of the token
func (a *Application) RateLimitFunc(token *Token, f func(int) int) error {
	var rateString string
	rateKey := serializeRateLimitKey(token.String())

	rateString, _ = a.store.GetKey(rateKey)
	if rateString == "" {
		rateString = "0"
	}

	rate, err := strconv.Atoi(rateString)
	if err != nil {
		return err
	}

	_, err = a.store.RemoveKey(rateKey)
	if err != nil {
		return err
	}

	var newRate = f(rate)

	if newRate < -1 {
		newRate = 0
	}

	return a.store.SetKey(rateKey, strconv.Itoa(newRate), token.expiresAt)
}

// GetRateLimit retrieve the Rate limit of the token
func (a *Application) GetRateLimit(token *Token) (int, error) {
	rateString, err := a.store.GetKey(serializeRateLimitKey(token.String()))

	rate, err := strconv.Atoi(rateString)
	if err != nil {
		return 0, err
	}

	return rate, err
}

func serializeTokenKey(token string) string {
	return token + tokenKeySuffix
}

func serializeRateLimitKey(token string) string {
	return token + rateLimitKeySuffix
}

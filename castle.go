package castle

import (
	"github.com/google/uuid"
	"github.com/sundowndev/castle/store"
	"time"
)

// Local store is an in-memory KV store for testing
type LocalStore = store.LocalStore

// Application defines an entry point for your web application
type Application struct {
	store      store.Store
	namespaces map[string]*Namespace
	scopes     map[string]*Scope
}

const UNLIMITED_RATE_LIMT Rate = -1
const RATE_LIMIT_KEY_SUFFIX string = ":rate"

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

	//a.SetRateLimit(func(_ Rate) Rate {
	//	return UNLIMITED_RATE_LIMT
	//})

	json, err := t.Serialize()
	if err != nil {
		return nil, err
	}

	err = a.store.SetKey(t.String(), json, expiration)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetToken retrieve a token from its value
func (a *Application) GetToken(token string) (*Token, error) {
	json, err := a.store.GetKey(token)
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
	_, err := a.store.RemoveKey(token)

	return err
}

// SetRateLimit mutate the Rate limit of the token
//func (a *Application) SetRateLimit(cb func (Rate) Rate) error {
//	//t.RateLimit = cb(t.RateLimit)
//}

// GetRateLimit retrieve the Rate limit of the token
//func (a *Application) GetRateLimit(token *Token) (Rate,error) {
//	rate, err = store.GetKey(t.String() + RATE_LIMIT_KEY_SUFFIX)
//}

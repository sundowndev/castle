package castle

import (
	"github.com/google/uuid"
	"github.com/sundowndev/castle/store"
)

type LocalStore = store.LocalStore

// Application defines an entry point for your web application
type Application struct {
	store      store.Store
	namespaces map[string]*Namespace
	scopes     map[string]*Scope
}

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

// NewToken creates a new UUID token inside
func (a *Application) NewToken(name string, expiration int, scopes ...*Scope) (*Token, error) {
	var scopesAsString = []string{}

	for _, v := range scopes {
		scopesAsString = append(scopesAsString, v.String())
	}

	t := &Token{
		uuid:      uuid.New(),
		Name:      name,
		Namespace: scopes[0].namespace.name,
		Scopes:    scopesAsString,
	}

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

// TODO: implement
// RevokeToken
func (a *Application) RevokeToken(token string) error {
	return nil
}

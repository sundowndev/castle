package castle

import (
	assertion "github.com/stretchr/testify/assert"
	"github.com/sundowndev/castle/store"
	"testing"
)

func TestNamespace(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should create scope from namespace", func(t *testing.T) {
		app := NewApp(store.NewLocalStore())

		ns := app.NewNamespace("test")

		scope := ns.NewScope("read")

		assert.Equal("test.read", scope.String())
		assert.Equal(ns, scope.namespace)
	})

	//t.Run("should create namespace from another namespace", func(t *testing.T) {
	//	app := NewApp(store.NewLocalStore())
	//
	//	ns1 := app.NewNamespace("ns1")
	//	ns2 := ns1.NewNamespace("ns2")
	//
	//	//scope := ns2.NewScope("read")
	//
	//	//assert.Equal(app, ns2.app)
	//	assert.Equal("ns1.ns2", ns2.name)
	//
	//	//assert.Equal("ns1.ns2.read", scope.String())
	//})
}

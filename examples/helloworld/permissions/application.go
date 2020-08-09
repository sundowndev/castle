package permissions

import (
	"github.com/sundowndev/castle"
	"github.com/sundowndev/castle/store"
)

var App *castle.Application

func init() {
	App = castle.NewApp(store.NewLocalStore())
}

package permissions

import "github.com/sundowndev/castle"

var Main *castle.Namespace

func init() {
	Main = App.NewNamespace("main")
}

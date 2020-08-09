package permissions

import "github.com/sundowndev/castle"

var Read *castle.Scope
var Write *castle.Scope
var Delete *castle.Scope

func init() {
	Read = Main.NewScope("read")
	Write = Main.NewScope("write")
	Delete = Main.NewScope("delete")
}

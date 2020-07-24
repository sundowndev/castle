package main

import (
	"fmt"
	"github.com/sundowndev/castle"
)

var App *castle.Application

var Repositories *castle.Namespace

var Read *castle.Scope

func init() {
	App = castle.NewApp(&castle.LocalStore{Store: make(map[string]string)})

	Repositories = App.NewNamespace("repositories")

	Read = Repositories.NewScope("read_repository")
}

func main() {
	token, _ := App.NewToken("myrepo", 3600, Read)

	json, _ := token.Serialize()

	fmt.Println(token.String(), json)
}

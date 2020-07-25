package main

import (
	"fmt"
	"github.com/sundowndev/castle"
	"time"
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
	token, _ := App.NewToken("myrepo", time.Now().Add(2 * time.Second), Read)

	json, _ := token.Serialize()

	fmt.Println(token.String(), json)

	token2, _ := App.GetToken(token.String())

	fmt.Println(token2.String(), token2.Name)
}

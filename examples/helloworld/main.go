package main

import (
	"fmt"
	"github.com/sundowndev/castle/examples/helloworld/permissions"
	"time"
)

func main() {
	token, _ := permissions.App.NewToken("hello", time.Now().Add(5*time.Minute), permissions.Read)

	fmt.Printf("Token: %v\n", token.String())
	fmt.Printf("Token has Read scope: %v\n", token.HasScope(permissions.Read))
	fmt.Printf("Token has Write scope: %v\n", token.HasScope(permissions.Write))

	json, _ := token.Serialize()

	fmt.Printf("Serialized token: %s", json)
}

package main

import (
	"github.com/sundowndev/castle"
	"github.com/sundowndev/castle/store"
	"net/http"
	"time"
)

var App *castle.Application

var Repositories *castle.Namespace

var Read *castle.Scope
var Write *castle.Scope

// Init application, namespaces and scopes
func init() {
	App = castle.NewApp(store.NewLocalStore())

	Repositories = App.NewNamespace("repositories")

	Read = Repositories.NewScope("read_repository")
	Write = Repositories.NewScope("write_repository")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token, err := App.NewToken("myrepo", time.Now().Add(2*time.Second), Write, Read)
		if err != nil {
			w.Write([]byte("error"))
			w.WriteHeader(500)
			return
		}

		w.Write([]byte(token.String()))
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		token, err := App.GetToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Write([]byte("error"))
			w.WriteHeader(500)
			return
		}

		rateLimit, _ := App.GetRateLimit(token)

		if !token.HasScope(Read) || rateLimit == 0 {
			w.WriteHeader(403)
			return
		}

		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":80", nil)
}

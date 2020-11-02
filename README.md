# castle

<div align="left">
  <a href="https://godoc.org/github.com/sundowndev/castle">
    <img src="https://godoc.org/github.com/sundowndev/castle?status.svg" alt="GoDoc">
  </a>
  <a href="https://github.com/sundowndev/castle/actions">
    <img src="https://img.shields.io/endpoint.svg?url=https://actions-badge.atrox.dev/sundowndev/castle/badge?ref=master" alt="build status" />
  </a>
  <a href="https://goreportcard.com/report/github.com/sundowndev/castle">
    <img src="https://goreportcard.com/badge/github.com/sundowndev/castle" alt="go report" />
  </a>
  <a href="https://codeclimate.com/github/sundowndev/castle/maintainability">
    <img src="https://api.codeclimate.com/v1/badges/e827d7cc994c6519d319/maintainability" />
  </a>
  <a href="https://codecov.io/gh/sundowndev/castle">
    <img src="https://codecov.io/gh/sundowndev/castle/branch/master/graph/badge.svg" />
  </a>
  <a href="https://github.com/sundowndev/castle/releases">
    <img src="https://img.shields.io/github/release/SundownDEV/castle.svg" alt="Latest version" />
  </a>
</div>

Access token management library for Go, backed by Redis. Designed for web services that needs a rate and time limited access control feature.

## Table of content

- [Design](#design)
- [Current status](#current-status)
- [Installation](#installation)
- [Usage](#usage)
    - [Stores](#stores)
    - [Namespaces](#namespaces)
    - [Scopes](#scopes)
    - [Using rate limit](#using-rate-limit)
- [Acknowledgement](#acknowledgement)
- [License](#license)

## Design

**Definitions** :

- **Application** : An entry point for your web service to register your store, namespaces and scopes.
- **Namespace** : Refers to a resource of your application.
- **Scope** : A permission of a namespace that can be granted to tokens (e.g: read, write, delete...).
- **Store**: A key/value storage system to store serialized tokens (Redis, etcd, RocksDB, MemcacheDB...).

**Principles** :

- Token value is RFC-4112 compliant.
- Token has a name, a rate limit, an expiration date and several scopes.
- Token's value and scopes **cannot be edited or altered**.
- Once created, if the token's lost, **it cannot be found anymore**.
- Uses KV store as the only source of truth.
- Token's rate limit cannot be lower than `0`. Default value `-1` is reserved to unlimited rate limit.

**Limitations** :

- **User feature compatibility** : This library does not fully support the user feature. The only way to make the user able to manage its created tokens is to store them in another database, which is an anti-pattern. We want Redis (or any store used) to be the only source of truth for credentials. Feel free to make design proposals.

## Current status

**Under active development**. The current version is v0, the API is unstable but still usable. The current design needs more feedback and use case examples to release a v1.

## Installation

```shell
go get -v -u github.com/sundowndev/castle
```

## Usage

### Stores

The default and recommended store is Redis. You can use it that way :

```go
package main

import (
  "github.com/sundowndev/castle"
)

func init() {
  _ = castle.NewApp(&castle.RedisStore{
    Host: "localhost",
    Port: 6739,
    Password: "",
    DB: "0",
  })
}
```

If you want to store tokens in-memory for testing, there's also a local store. LocalStore uses go routines to revoke expired tokens. So it should have the exact same behavior as Redis.

```go
package main

import (
  "github.com/sundowndev/castle"
  "github.com/sundowndev/castle/store"
)

func init() {
  _ = castle.NewApp(store.NewLocalStore())
}
```

#### Creating your own store

You can create your own store by using the following `Store` interface :

```go
type Store interface {
	GetKey(string) (string, error)
	SetKey(string, string, time.Time) error
	RemoveKey(string) (bool, error)
	Flush() error
}
```

Ideally, a store would follow the following design principles :

- A key cannot be updated, it must be removed first
- Store must have a lock system to avoid conflict in go routines
- Store never panics, it always returns errors cleanly

### Namespaces

Define some namespaces.

```go
package main

import (
  "github.com/sundowndev/castle"
)

var Repositories *castle.Namespace

func init() {
  Repositories = App.NewNamespace("repositories")
}
```

### Scopes

Define some scopes inside namespace.

```go
package main

import (
  "github.com/sundowndev/castle"
)

var Repositories *castle.Namespace

var Read *castle.Scope
var Write *castle.Scope

func init() {
  Repositories = App.NewNamespace("repositories")
  
  Read = Repositories.NewScope("read") // repositories.read
  Write = Repositories.NewScope("write") // repositories.write
}
```

### Using rate limit

Rate limit feature is optional. By default, rate limit is set to `-1`, which basically means the token has no rate limit. You're free to handle this feature your own way, decreasing or increasing the rate limit as you want. You may also use a worker service to periodically reset the rate limit of tokens.

```go
package controllers

func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
    token, err := app.NewToken("token_name", time.Now().Add(1 * time.Minute), read)
    err != nil {
        // Handle err...
        return
    }

    // set rate limit to 1000
    App.RateLimitFunc(token, func (rate int) int {
        return 1000
    })

    // return token to client...   
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
    token, err := app.GetToken(r.Header.Get("Authorization"))
    err != nil {
        // Handle err...        
    }

    if rate, _ := App.GetRateLimit(token); rate == 0 {
        w.WriteHeader(403)
        return
    }

    // set rate limit to 999
    App.RateLimitFunc(token, func (rate int) int {
        return rate - 1
    })

    // read resource...
}
```

## Acknowledgement

- [node-roles](https://dresende.github.io/node-roles/) (Node)
- [kan](https://github.com/davydovanton/kan) (Ruby)
- [rolify](https://github.com/RolifyCommunity/rolify) (Ruby)
- [Laravel Auth](https://github.com/jeremykenedy/laravel-auth) (PHP)
- [Gitlab's access token feature](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)

## License

This project is licensed under the [GPL-3.0 License](LICENSE).

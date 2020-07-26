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

# castle

Access token management backed by Redis. Designed for large scale systems with several permissions in different contexts (e.g: Gitlab, GitHub...), but also simpler systems (e.g: Nextcloud, HaveIBeenPwned...).

## Table of content

- [Design principles](#design-principles)
- [Current status](#current-status)
- [Installation](#installation)
- [Usage](#usage)
    - [Stores](#stores)
    - [Namespaces](#namespaces)
    - [Scopes](#scopes)
    - [Using rate limit](#using-rate-limit)
- [Acknowledgement](#acknowledgement)
- [License](#license)

## Design principles

**Definitions :**

- **Application** : An entry point for your web service to register your store, namespaces and scopes.
- **Namespace** : Refers to a resource of your application.
- **Scope** : A permission of a namespace that can be granted to tokens.
- **Store**: A key/value storage system to store serialized tokens.

**Principles** :

- Token value is RFC-4112 compliant
- Token has a name, a single namespace, a rate limit and several scopes
- Tokens **cannot be permanent, edited or altered**
- Tokens cannot be gathered in mass through the API
- Once created, if the token is lost, **it cannot be found anymore**

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
)

func init() {
  _ = castle.NewApp(&castle.LocalStore{Store: map[string]string{}})
}
```

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

...

### Using rate limit

Rate limit feature is optional. By default, rate limit is set to `-1`, which basically means the token has no rate limit. You're free to handle this feature your own way. Decreasing or increasing the rate limit as you want. You may also use a worker service to periodically reset the rate limit of tokens.

```go
package controllers

func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
    token, err := app.NewToken("token_name", time.Now().Add(1 * time.Minute), read)
    err != nil {
        // Handle err...        
    }

    token.SetRateLimit(500)

    // return token to client...   
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
    token, err := app.GetToken(r.Header.Get("Authorization"))
    err != nil {
        // Handle err...        
    }

    if token.RateLimit == 0 {
        w.WriteHeader(403)
    }

    token.RateLimit(func (rate int) int {
        return rate - 1
    })

    // return token to client...   
}
```

## Acknowledgement

- [node-roles](https://dresende.github.io/node-roles/) (Node)
- [kan](https://github.com/davydovanton/kan) (Ruby)
- [rolify](https://github.com/RolifyCommunity/rolify) (Ruby)
- [Laravel Auth](https://github.com/jeremykenedy/laravel-auth) (PHP)

## License

This project is licensed under the [GPL-3.0 License](LICENSE).

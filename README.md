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

Access token management backed by Redis. Designed for APIs and micro services. Written for large scale systems with several permissions and roles in different contexts (e.g: Gitlab, GitHub...), but also simplier systems (e.g: Nextcloud, HaveIBeenPwned...).

## Table of content

- [Design principles](#design-principles)
- [Current status](#current-status)
- [Installation](#installation)
- [Usage](#usage)
- [Acknowledgement](#acknowledgement)
- [License](#license)

## Design principles

**Definitions :**

- **Applications** : ...
- **Namespace** : ...
- **Scope** : ...

**Principles** :

- Token value is RFC-4112 compliant
- Token has a name, a single namespace and several scopes
- Tokens **cannot be persistant, edited or altered**
- Tokens cannot be gathered in mass through the API
- Once created, if the token is lost, **it cannot be found anymore**

## Current status

The current version is v0, the API is instable but still usable. The current design needs more feedback and use case examples to release a v1.

## Installation

```shell
go get github.com/sundowndev/castle
```

## Usage

First, define your application :

```go
package main

import (
  "github.com/sundowndev/castle"
)

var App *castle.Application

func init() {
  App := castle.NewApp(&castle.RedisStore{
    Host: "localhost",
    Port: 6739,
    Password: "",
    DB: "0",
  })
}
```

then define some namespaces

```go
package main

import (
  "github.com/sundowndev/castle"
)

var Repositories *castle.Namespace

func init() {
  Repositories := App.NewNamespace("repositories")
}
```

## Acknowledgement

- [node-roles](https://dresende.github.io/node-roles/) (Node)
- [kan](https://github.com/davydovanton/kan) (Ruby)
- [rolify](https://github.com/RolifyCommunity/rolify) (Ruby)
- [Laravel Auth](https://github.com/jeremykenedy/laravel-auth) (PHP)

## License

This project is licensed under the [GPL-3.0 License](LICENSE).

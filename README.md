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

A role management library for Go. Supports both static and dynamic role assignment, so your design shouldn't be impacted. Written for large scale systems with several permissions and roles in different contexts (e.g: Gitlab, GitHub, ...), but also simplier systems (e.g: Nextcloud, ...).

## Table of content

- [Design principles](#design-principles)
- [Current status](#current-status)
- [Installation](#installation)
- [Usage](#usage)
- [Acknowledgement](#acknowledgement)
- [License](#license)

## Design principles

**Definitions :**

- **Applications** : An application is a set of roles. You can create multiple applications with different roles.
- **Permissions** : A permission, but is nothing without a role assigned to it. Permissions can be shared between applications.
- **Roles** : A role is a set of permissions.
<!--
- **Abilities** : ...

**Principles** :

- A profile can have only one role per context, but can have many roles from many contexts -->

## Current status

The current version is v0, the API is instable but still usable. The current design needs more feedback and use case examples to release a v1.

## Installation

```shell
go get github.com/sundowndev/castle
```

## Usage

First, define your application. The name must match `([a-zA-Z])\w+`.

```go
package main

import (
  "github.com/sundowndev/castle"
)

var App *castle.Application

func init() {
  App, err = castle.NewApplication("myapp")

  if err != nil {
    panic(err) // Validation error
  }
}
```

Define some permissions :

```go
package permissions

import (
  "myapp"
  "github.com/sundowndev/castle"
)

var DeleteAnyVideo *castle.Permission
var UploadVideo *castle.Permission

func init() {
  DeleteAnyVideo = myapp.App.NewPermission()
  UploadVideo = myapp.App.NewPermission()
}
```

Define some roles :

```go
package roles

import (
  "myapp"
  "myapp/roles"
  "github.com/sundowndev/castle"
)

var Admin *castle.Role
var Guest *castle.Role

func init() {  
  // Assign permissions to roles
  // Note returned error was ignored in this example
  Guest, _ = myapp.App.NewRole("guest", roles.UploadVideo)
  Admin, _ = myapp.App.NewRole("admin", roles.DeleteAnyVideo).InheritFrom(Guest) // Admin role will inherit from Guest's permissions
}
```

Check a role's permissions :

```go
package main

import (
  "myapp"
  "myapp/roles"
  "github.com/sundowndev/castle"
)

func main() {
  role, err := myapp.App.GetRole("myapp.admin")

  if err != nil {
    panic(err) // This role doesn't exists
  }

  if true != role.HasPermission(roles.UploadVideo) {
    // Handle err
  }

  // Admin role has UploadVideo role/permission
}
```

## Acknowledgement

- [node-roles](https://dresende.github.io/node-roles/) (Node)
- [kan](https://github.com/davydovanton/kan) (Ruby)
- [rolify](https://github.com/RolifyCommunity/rolify) (Ruby)
- [Laravel Auth](https://github.com/jeremykenedy/laravel-auth) (PHP)

## License

This project is licensed under the [GPL-3.0 License](LICENSE).

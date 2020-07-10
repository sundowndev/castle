# castle

A role management library for Go.

## Background

...

## Current status

v0, not stable...

## Definitions

- Application : ...
- Role : ...
- Profile : ...

## Installation

```
$ go get github.com/sundowndev/castle
```

## Usage

Define your application :

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

Define some roles :

```go
package roles

import (
  "myapp"
  "github.com/sundowndev/castle"
)

var DeleteAnyVideo *castle.Role
var UploadVideo *castle.Role

func init() {
  // Create some roles
  DeleteAnyVideo = myapp.App.NewRole()
  UploadVideo = myapp.App.NewRole()
}
```

Define some profiles :

```go
package profiles

import (
  "myapp"
  "myapp/roles"
  "github.com/sundowndev/castle"
)

var Admin *castle.Profile
var Guest *castle.Profile

func init() {  
  // Assign roles to profiles
  // Note returned error was ignored in this example
  Guest, _ = myapp.App.NewProfile("guest", roles.UploadVideo)
  Admin, _ = myapp.App.NewProfile("admin", roles.DeleteAnyVideo).InheritFromProfile(Guest) // Admin profile will inherit from Guest's permissions
}
```

Check a profile's permissions :

```go
import (
  "myapp"
  "myapp/roles"
  "github.com/sundowndev/castle"
)

func main() {
  profile, err := myapp.App.GetProfile("admin")

  if err != nil {
    panic(err) // This profile doesn't exists
  }

  if true != profile.HasRole(roles.UploadVideo) {
    // Handle err
  }

  // Admin profile has UploadVideo role/permission
}
```

## Database integrity

...

## Acknowledgement

- [node-roles](https://dresende.github.io/node-roles/)

## License

GPL v3.

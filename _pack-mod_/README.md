# Packages and Modules

## Modules

### Initialize a Module

```sh
go mod init <domain.fqdn/namespace/module-name
```

### Retrive Dependencies

```sh
go get -u <domain.fqdn/namespace/module-name>
```

### List Dependencies

```sh
go list
go list -m all
go list -v all
go list -m versions <module-path>
```

### Cleanup

```sh
go mod verify
go mod tidy
```

### Bundle dependencies locally

```sh
go mod vendor
```

...used with `go build/run`

```sh
go build -mod=vendor
go run -mod=vendor
```

## Packages

### Library packages (`package services`)

- Contains only code that is used by other packages.
- Are not intended to be used by the end user.
- Name must match the directory name.


### Main packages (`package main`)

- Application entry point.
- Contains a `main()` function.
- Can be in any directory.
- Focus in app setup and initialization.


```sh
go run _pack-mod_/main.go
```

### Lifecycle
Import -> Set Variables -> Init()
Multiple init() functions can be called.

### Member visibility

Public:
- Capitalized
- Visible to other packages.

Package:
- Lowercase
- Visible within the package.

Internal:
- Scoped to the package and its descendants.
- Can use public and package -level members.


### Documention a package
```sh
go doc github.com/dragosh/go-learn/_pack-mod_/mypkg/
go doc github.com/dragosh/go-learn/_pack-mod_/mypkg/ StartServer
```

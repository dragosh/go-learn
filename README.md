# GO Learn

## Topics:

- [x] Debugging
- [] Testing

### Debugging /w dlv

Essentials debugging

1. Install

Get it from https://github.com/go-delve/delve/tree/master/Documentation/installation

2. Start debugger

```sh
dlv debug pkg/debug/main.go

help
```

3. Breakpoint

```sh
b main.main // set breakpoint in main fn
b main.go:15 // set breakpoint at line
b main.go:19 // set breakpoint
```

4. Continue and locals

```sh
c // continue
p hello // print var
c // continue
locals // print local vars
p hello
```

### Testing

1. Single file

```sh
go test pkg/testing/sample_test.go
```

2. Test the package

```sh
cd pkg/testing
go test # run all
go test -run Greet # only regexp
go test -cover # cover reporting terminal
go test -coverprofile cover.out # cover reporting per test
go tool cover -func cover.out # cover reporting function
go tool cover -html cover.out # cover reporting html
go test -coverprofile count.out -covermode count  # cover reporting html count
go test -v # Show running tests
```


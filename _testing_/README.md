## Testing

### Single file

If `main_test.go` and `main.go` are the same package (a common case) then you must name all other files required to build `main_test`:

```sh
go test _testing_/main_test.go _testing_/main.go
```

### Folder and pattern test (Greet) specific

`-count 1` also ensures that the test is ran every time instead of being cached. Useful when you are testing against race conditions and have a test that fails only sometimes.

```sh
go test -count 1 -run Greet ./_testing_/
```

1. Test the package

```sh
cd _testing_
go test # run all
go test -run Greet # only regexp
go test -cover # cover reporting terminal
go test -coverprofile cover.out # cover reporting per test
go tool cover -func cover.out # cover reporting function
go tool cover -html cover.out # cover reporting html
go test -coverprofile count.out -covermode count  # cover reporting html count
go test -v # Show running tests
```

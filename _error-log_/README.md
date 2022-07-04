# Error and Logging

```sh
go run _error-log_/main.go
```

```sh
#show only logging output
go run _error-log_/main.go  > /dev/null
```



## Error type is an interface type

```go
type error interface {
    Error() string
}
```

# Data types

- Collections (Arrays, Slices, Maps)
- Records (Structs)

```ssh
go run _data-types_/main.go
```


## Arrays

```go
// varname := [size]type{default, values}
array = [3]int{1, 2, 3}

// varname []type{default, values}
dynArray = []int{1, 2, 3}

```


## Slices

```go
// varname := from[range] // make a copy
slice := arr[:] // ponting to the array NOT immutable

// varname make([]type,len,cap)
slice := make([]string, 0, 3)
```

## Maps

```go
mapName := map[string]int{"magic": 42}

m := make(map[string]int)
```

## Structs

```go
  u3 := struct {
    ID        int
    FirstName string
    LastName  string
  }{1, "first", "last"}
```


## Debugging /w dlv

### Essentials debugging

1. Install

Get it from https://github.com/go-delve/delve/tree/master/Documentation/installation

2. Start debugger (from terminal)

```sh
dlv debug github.com/dragosh/go-learn/_debugging_

help
```

1. Breakpoint

```sh
b main.main // set breakpoint in main fn
b main.go:21 // set breakpoint at line
b main.go:24 // set breakpoint at line
```

4. Continue and locals

```sh
c // continue
p hello // print var
c // continue
locals // print local vars
p hello
```

5. Debug commands file

```sh
dlv debug github.com/dragosh/go-learn/_debugging_
source pkg/debug/debug-cmd.txt
```

5. Frames / goroutine

```sh
dlv debug github.com/dragosh/go-learn/_debugging_
frame <index> locals
```

5. Inspection

```
print <expr>
var <regexp>
funcs <regexp>
types <regexp>
sources <regexp>
sources go-learn # regexp
list

disassemble (assembly output)
regs (CPU registers)
```

VSCode setup https://github.com/vscode-debug-specs/go

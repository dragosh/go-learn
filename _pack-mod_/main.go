/*
* -----------------------------------------------------------
* Packages and modules
* -----------------------------------------------------------
 */
package main

import (
	"fmt"

	"github.com/dragosh/go-learn/_pack-mod_/mypkg"
)

func main() {
	mypkg.StartServer()
	mypkg.OtherMethod()
}

func init() {
	fmt.Println("Init in main.go")
}

/*
* -----------------------------------------------------------
* Debuging
* -----------------------------------------------------------
 */

package main

import (
	"fmt"
)

var hello string = "Hello"

func sayHello() {
	fmt.Println("Say: ", hello)
}

func main() {
	// add a breakpoint on the next line
	hello = "Hi"

	magic := 42
	// add a breakpoint on the next line
	sayHello()

	fmt.Println(magic)

}

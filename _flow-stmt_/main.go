/*
* -----------------------------------------------------------
* Flow and Statements (loops, panics, ifs, switch)
* -----------------------------------------------------------
 */

package main

import (
	"fmt"
	"reflect"
)

func forStm() {
	/*
	* -----------------------------------------------------------
	* For
	* -----------------------------------------------------------
	 */
	var i int   //explicit
	for i < 5 { // loop till condition exit
		fmt.Println(i)
		i++
		if i == 3 {
			break // exit completely
			// continue // skip on 3
		}
		fmt.Println("continuing...")
	}

	for i := 0; i < 5; i++ { // implicit with scoped variable i
		fmt.Println(i)
	}

	for { // infinite loop
		if i < 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	slice := []int{1, 2, 3}
	for i, v := range slice { // similar to js foreach
		fmt.Println(i, v)
	}

	httpPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range httpPorts { // similar to js map
		fmt.Println(k, v)
	}

	for _, v := range httpPorts { // only the value
		fmt.Println(v)
	}
}

func ifStm() {
	/*
	* -----------------------------------------------------------
	* Ifs
	* -----------------------------------------------------------
	 */
	if 1 == 2 {
		// same
	} else if reflect.TypeOf(1).Kind() == reflect.TypeOf(2).Kind() {
		// same types
		fmt.Println("Same type")
	} else {
		// different
	}

	//A statement can precede conditionals;
	// any variables declared in this statement are available in all branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switchStm() {
	/*
	* -----------------------------------------------------------
	* Switch
	* -----------------------------------------------------------
	 */

	type HTTPRequest struct {
		Method string
	}

	req := HTTPRequest{Method: "GET"}

	switch req.Method {
	case "GET":
		fmt.Println("Get request") // implicit break
		//fallthrough            //go down through cases
	case "POST":
		fmt.Println("Post request")
	default:
		fmt.Println("Unkown request type")
	}
}

/*
* -----------------------------------------------------------
* panic -> similar with Exceptions
* Indicates an exceptional condition that there's no easy way to recover from
* -----------------------------------------------------------
 */
func panicStm() {

	// `recover` stops the panic and returns the value that was passed to the call to panic.
	//	We have to pair it with defer
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Something bad happend") // exit code 1

}

// `defer`` is often used when resources need to be freed in some way.
// For example when we open a file we need to make sure to close it later.
func deferControl() {
	first := func() {
		fmt.Println("First.")
	}
	second := func() {
		fmt.Println("Second")
	}
	defer second()
	first()
}

func main() {

	forStm()
	ifStm()
	switchStm()
	deferControl()

	panicStm()
}

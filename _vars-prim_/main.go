/*
* -----------------------------------------------------------
* Variables & Primitives
* -----------------------------------------------------------
 */

package main

import "fmt"

const PI = 3.14 // global constant

const (
	IMPORT_CONST_ONE = "ONE"
	IMPORT_CONST_TWO = "TWO"
	// https://www.dannyvanheumen.nl/post/when-to-use-go-iota/
	// Use iota for purpose of “internal” use only. Use iota assignments if the value does not really matter.
	IMPORT_CONST_IOTA   = iota
	IMPORT_CONST_IOTA_3 = iota
)

func primitives() {
	/*
	* -----------------------------------------------------------
	* Primitives
	* -----------------------------------------------------------
	 */
	var explicit_int int = 42    // explicit int and all the range `float` `bigint` etc.
	implicit_string := "Dragosh" // implicit string
	boolean := true              // bool
	comp := complex(4, 2)        // complex (real and imaginary) used for math
	real, imag := real(comp), imag(comp)
	fmt.Println(explicit_int)
	fmt.Println(implicit_string)
	fmt.Println(boolean)
	fmt.Println(comp)
	fmt.Println(real, imag)
}

func pointers() {
	/*
	* -----------------------------------------------------------
	* Pointers
	* -----------------------------------------------------------
	 */
	var expPointName *string = new(string) // explicit assignment
	*expPointName = "Dragosh"
	fmt.Println(*expPointName)

	impPointName := "Dragosh"
	fmt.Println(impPointName)

	ptr := &impPointName
	fmt.Println(ptr, *ptr)

	impPointName = "Florin"
	fmt.Println(ptr, *ptr)
}
func operations() {
	fmt.Println("go" + "lang")        // concatenate 2 strings
	fmt.Println("1+1 =", 1+1)         // add 2 int
	fmt.Println("7.0/3.0 =", 7.0/3.0) // divide 2 floats
	// bool evaluation
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func constants() {
	/*
	* -----------------------------------------------------------
	* Constants
	* -----------------------------------------------------------
	 */

	const cons int = 42 // implicit
	const pi = 3.1415   // explicit
	fmt.Println(cons + 8)
	fmt.Println(pi)
	fmt.Println(float32(cons) + 0.2) //need to convert the type
	fmt.Println(pi + 4)              // this will work
	fmt.Println(PI)                  // global
	fmt.Println(IMPORT_CONST_ONE, IMPORT_CONST_TWO, IMPORT_CONST_IOTA, IMPORT_CONST_IOTA_3)
}
func main() {

	operations()
	primitives()
	constants()
	pointers()

}

/*
* -----------------------------------------------------------
* Methods & Interfaces
* -----------------------------------------------------------
* A method is a function with a special receiver argument.
* The receiver appears in its own argument list between the func keyword and the method name.
*
* An interface type is defined as a set of method signatures or
* Interfaces are named collections of method signatures.
 */
package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float64
}

/*
* -----------------------------------------------------------
* Method
* -----------------------------------------------------------
 */

// func <reciver type> name(arg type) return type

// Methods can be defined for either pointer or value receiver types.
func (r *rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

/*
* -----------------------------------------------------------
* Interface
* -----------------------------------------------------------
 */
type geometry interface {
	area() float64
	perim() float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	//Go automatically handles conversion between values and pointers for method calls.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	c := circle{radius: 5}

	measure(rp)
	measure(c)

	// Empty interfaces are used by code that handles values of unknown type.
	var empty interface{}
	empty = 42
	fmt.Printf("(%v, %T)\n", empty, empty)
	empty = "hello"
	fmt.Printf("(%v, %T)\n", empty, empty)
}

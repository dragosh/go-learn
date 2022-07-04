/*
* -----------------------------------------------------------
* Testing
* -----------------------------------------------------------
 */

package main

import "fmt"

// Public
func Greet(name string) string {
	return fmt.Sprintf("Hello %v!\n", name)
}

// Private
func depart(name string) string {
	return fmt.Sprintf("Good bye %v!\n", name)
}

func main() {
	fmt.Print(Greet("Dragosh"))
	fmt.Print(depart("Dragosh"))
}

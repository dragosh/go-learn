/*
* -----------------------------------------------------------
* Functions
* -----------------------------------------------------------
 */
package main

import "fmt"

// that takes two ints and returns their sum as an int.
func simplePlus(a int, b int) int {
	return a + b
}

// you may omit the type name for the like-typed parameters
// up to the final parameter that declares the type.
func omitType(a, b, c int) int {
	return a + b + c
}

//  `(int, int) `in this function signature shows that the function returns 2 ints.
func returnMultipleValues() (int, int) {
	return 3, 7
}

// can be called with any number of trailing arguments
func variadicSum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// anonymous functions, which can form closures.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// nested functions
func nested(str string) {
	func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function")
	}(str)

	funcVar := func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function assigned to a variable.")
	}

	funcVar(str)
}

// callback functions
func startCallback(str string, callback func(string)) {
	callback(str)
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	/*
	* -----------------------------------------------------------
	* Simple
	* -----------------------------------------------------------
	 */
	res := simplePlus(1, 2)
	fmt.Println("1+2 =", res)
	res = omitType(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	/*
	* -----------------------------------------------------------
	* Return multiple values
	* -----------------------------------------------------------
	 */

	a, b := returnMultipleValues()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := returnMultipleValues()
	fmt.Println(c)
	//fmt.Println(_) //throws an error

	/*
	* -----------------------------------------------------------
	* Variadic args
	* -----------------------------------------------------------
	 */
	variadicSum(1, 2)
	variadicSum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	// apply a slice to a variadic function using func(slice...)
	variadicSum(nums...)

	/*
	* -----------------------------------------------------------
	* Closures
	* -----------------------------------------------------------
	 */
	// This function value captures its own i value,
	// which will be updated each time we call nextInt
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	/*
	* -----------------------------------------------------------
	* Nested
	* -----------------------------------------------------------
	 */

	nested("Dragosh")
	/*
	* -----------------------------------------------------------
	* Callback
	* -----------------------------------------------------------
	 */
	startCallback("Dragosh", func(str string) {
		fmt.Printf("Hi %v \n", str)
	})

	/*
	* -----------------------------------------------------------
	* Recursion
	* -----------------------------------------------------------
	 */
	fmt.Println("Factorial of (3) =", fact(3))
}

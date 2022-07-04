/*
* -----------------------------------------------------------
* Data Types
* -----------------------------------------------------------
 */

package main

import (
	"fmt"
	"reflect"
)

func arrays() {
	/*
	* -----------------------------------------------------------
	* Arrays
	* Fixed sized collection of similar data types
	* -----------------------------------------------------------
	 */

	// varname [size]type
	var arr [3]int // long syntax

	fmt.Println("empty:", arr)
	arr[0] = 1
	arr[2] = 3

	fmt.Println("all", arr)
	fmt.Println("one", arr[2])
	// varname [size]type{default, values}
	arr2 := [3]int{1, 2, 3} //short syntax
	fmt.Println(arr2)
	fmt.Println(arr2[2])

	fmt.Println("len:", len(arr2))
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func slices() {

	/*
	* -----------------------------------------------------------
	* Slices
	* NOT a fix sized array
	* Built on top of arrays
	* -----------------------------------------------------------
	 */
	arr := [3]int{1, 2, 3} // original array
	// varname := from[range] // make a copy
	slice := arr[:] // ponting to the array NOT immutable

	arr[2] = 41
	slice[2] = 42
	fmt.Println(arr, slice)

	// varname []type{default, values}
	dynslice := []int{1, 2, 3}
	fmt.Println(dynslice)
	// assign an array to fit data (copy operation)
	dynslice = append(dynslice, 4)
	fmt.Println(dynslice)
	// varname := from[range]
	s1 := dynslice[1:]
	s2 := dynslice[:2]
	s3 := dynslice[1:2]
	from := 1
	to := 4
	s4 := dynslice[from:to]
	fmt.Println(s1, s2, s3, s4)

	// using make
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	printSlice("s", s)
	// make a copy
	c := make([]string, len(s)) //without values
	copy(c, s)                  // add the values
	printSlice("c", c)

	// multi dim
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

func maps() {
	/*
		* -----------------------------------------------------------
		* Maps
		* A colletion that associates arbitrary keys to value
		* built-in associative data type
		 (sometimes called object, hashes or dicts in other languages).
		* -----------------------------------------------------------
	*/
	// mapname := map[key_type]value_type
	mapName := map[string]int{"magic": 42} // implicit declaration
	fmt.Println(mapName)
	mapName["magic"] = 7
	fmt.Println(mapName)

	// delete
	delete(mapName, "magic")
	fmt.Println(mapName)

	// using make
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map.
	_, hasKey := m["k2"]
	fmt.Println("HasKey k2:", hasKey)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func structs() {
	/*
	* -----------------------------------------------------------
	* Structs
	* We can associate any type of data
	* More like Custom type / Union type
	* typed collections of fields.
	* They’re useful for grouping data together to form records.
	* -----------------------------------------------------------
	 */
	type user struct {
		ID        int    //when empty, 0  is the default value
		FirstName string //when empty, '' is the default value
		LastName  string
	}
	var u user //explicit
	fmt.Println("Empty", u)

	u.ID = 1
	u.FirstName = "Dragosh"
	u.LastName = "OZ"

	fmt.Println("Explicit", u)
	fmt.Println("Single value", u.FirstName)

	u2 := user{ //implicit
		ID:        1,
		FirstName: "Dragosh",
		LastName:  "OZ",
	}
	fmt.Println("Implicit", u2)
	// define and assign
	u3 := struct {
		ID        int
		FirstName string
		LastName  string
	}{1, "first", "last"}

	fmt.Println("Def+Default", u3)

	// constructs
	fmt.Println(person{"Dragosh", 42})
	//named fields
	fmt.Println(person{name: "OZ"})
	jon := newPerson("Jon")
	fmt.Println(jon)
	fmt.Println(jon.age)

	sean := person{name: "Sean", age: 50}
	seanPointer := &sean
	fmt.Println(seanPointer.age)

	//An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

}

func iterate() {
	/*
	* -----------------------------------------------------------
	* `range` iterates over elements in a variety of data structures.
	* -----------------------------------------------------------
	 */

	nums := []int{2, 3, 4}
	sum := 0

	// Iterate over slices/arrays
	for i, num := range nums {
		fmt.Println("Index:", i, "Value", num)
		sum += num
	}
	fmt.Println("Total sum:", sum)

	// Iterate over maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("key:", k)
		fmt.Printf("%s -> %s\n", k, v)
	}
	// Iterate over strings (runeValues)
	// https://gobyexample.com/strings-and-runes
	for idx, runeValue := range "go" {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Interate through structs
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
}
func main() {
	arrays()
	slices()
	maps()
	structs()
	iterate()
}

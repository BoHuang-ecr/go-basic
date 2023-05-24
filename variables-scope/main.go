/*
	The scope is the range in the source code of the constant, type, variable, function, or package represented by the declared identifier.
	Variables in Go language can be declared in three places:
		* Variables defined within a function are called local variables
		* Variables defined outside a function are called global variables
		* Variables in a function definition are called formal variables
*/

package main

import "fmt"

// global variable
var a int = 20

func main() {
	// local variable in main function
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("inside main() a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("insde main() c = %d\n", c)
}

// formal variable
func sum(a, b int) int {
	fmt.Printf("inside sum() a = %d\n", a)
	fmt.Printf("inside sum() b = %d\n", b)

	return a + b
}

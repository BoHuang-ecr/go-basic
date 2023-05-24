/*
Go language variable names are composed of letters, numbers, and underscores, and the first character cannot be a number.
The general form of declaring a variable is to use the var keyword:
*/

package main

import "fmt"

func main() {
	var a string = "hello world"
	fmt.Println(a)

	var b_int, c_int int = 1, 2
	fmt.Println(b_int, c_int)

	// define variable without init value.
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// var ff string = "string"
	ff := "string"
	fmt.Println(ff)
	fmt.Printf("The type of variable ff is: %T\n", ff)

	// pointer
	fmt.Printf("The address of variable a is: %p\n", &a)
	fmt.Printf("The value of variable a is: %s\n", *&a)
}

/*
	func function_name( [parameter list] ) [return_types] {
		body
	}

	func (variable_name variable_data_type) function_name() [return_type]{
   	    body
	}
*/

package main

import "fmt"

func main() {
	// local variable
	var a int = 100
	var b int = 200

	fmt.Printf("before swap, the value of a: %d\n", a)
	fmt.Printf("before swap, the value of b: %d\n", b)

	swap(&a, &b)

	fmt.Printf("after swap, the value of a: %d\n", a)
	fmt.Printf("after swap, the value of b: %d\n", b)

	// Go does not have classes. However, you can define methods on types.
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("The area of circle is = ", c1.getArea())
}

func swap(x *int, y *int) {
	/*
		Passing refers to copying the actual parameters to the function when calling the function,
		so that if the parameters are modified in the function, the actual parameters will not be affected.

		Passing by reference refers to passing the address of the actual parameter to the function when calling the function,
		then the modification of the parameter in the function will affect the actual parameter.
	*/
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// define struce
type Circle struct {
	radius float64
}

// this is method of data type Circle
func (c Circle) getArea() float64 {
	//c.radius
	return 3.14 * c.radius * c.radius
}

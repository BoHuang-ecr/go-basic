/*
	The built-in operators in Go language are:
		* arithmetic operator
		* relational operator
		* Logical Operators
		* bitwise operator
		* assignment operator
		* other operators
*/

package main

import "fmt"

func arithmetic() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("The first line - the value of c is %d \n ", c)
	c = a - b
	fmt.Printf("The second line - the value of c is % d \n ", c)
	c = a * b
	fmt.Printf("The third line - the value of c is %d \n ", c)
	c = a / b
	fmt.Printf("The fourth line - the value of c is %d \n ", c)
	c = a % b
	fmt.Printf("The fifth line - the value of c is %d \n ", c)
	a++
	fmt.Printf("The sixth line - the value of a is %d\n ", a)
	a = 21 // For the convenience of testing, a is reassigned to 21 here
	a--
	fmt.Printf("The seventh line - the value of a is %d \n ", a)
}

func relational() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("Line 1 - a is equal to b \n")
	} else {
		fmt.Printf("Line 1 - a is not equal to b \n")
	}
	if a < b {
		fmt.Printf("Second line - a is less than b \n")
	} else {
		fmt.Printf("Second line - a is not less than b \n")
	}
	if a > b {
		fmt.Printf(" Third line - a is greater than b \n")
	} else {
		fmt.Printf("Line 3 - a is not greater than b \n")
	}
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("Line 4 - a is less than or equal to b \n")
	}
	if b >= a {
		fmt.Printf("Line 5 - b is greater than or equal to a \n")
	}
}

func logical() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("Line 1 - condition is true \n")
	}
	if a || b {
		fmt.Printf("Line 2 - condition is true \n")
	}
	/* modify the values ​​of a and b */
	a = false
	b = true
	if a && b {
		fmt.Printf("Line 3 - condition is true \n")
	} else {
		fmt.Printf("Line 3 - condition is false \n")
	}
	if !(a && b) {
		fmt.Printf("Line 4 - condition is true \n")
	}
}

func bitwise() {

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("The first line - the value of c is %d \n ", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("The second line - the value of c is %d \n ", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("The third line - the value of c is %d \n ", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("The fourth line - the value of c is %d \n ", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - the value of c is %d \n ", c)
}

func assignment() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("Line 1 - = operator instance, value of c = %d \n ", c)

	c += a
	fmt.Printf("Line 2 - += operator instance, c value = %d \n ", c)

	c -= a
	fmt.Printf("Line 3 - -= operator instance, c value = %d \n ", c)

	c *= a
	fmt.Printf("Line 4 - *= operator instance, c value = %d \n ", c)

	c /= a
	fmt.Printf("line 5 - /= operator instance, c value = %d \n ", c)

	c = 200

	c <<= 2
	fmt.Printf("line Line 6 - <<= operator instance, c value = %d \n ", c)

	c >>= 2
	fmt.Printf("line 7 - >>= operator instance, c value = %d \n ", c)

	c &= 2
	fmt.Printf("line 8 - &= operation operator instance, c value = %d \n ", c)

	c ^= 2
	fmt.Printf("line 9 - ^= operator instance, c value = %d \n ", c)

	c |= 2
	fmt.Printf("Line 10 - |= operator instance, c value = % d\n", c)

}

func other() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int /* operator instance */
	fmt.Printf("Line 1 - a variable type is = %T \n ", a)
	fmt.Printf("Line 2 - variable type b is = %T \n ", b)
	fmt.Printf("Line 3 - variable type c is = %T\n ", c)

	/* & and * operator instances*/
	ptr = &a /* 'ptr' contains the address of variable 'a'*/
	fmt.Printf("a is %d \n ", a)
	fmt.Printf("*ptr is %d \n ", *ptr)
}

func main() {
	arithmetic()
	relational()
	logical()
	bitwise()
	assignment()
	other()
}

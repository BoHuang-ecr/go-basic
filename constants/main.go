/*
const identifier [type] = value

iota,
a special constant, can be considered as a constant that can be modified by the compiler.
iota will be reset to 0 when the const keyword appears (before the first line inside const),
and each new line of constant declaration in const will make iota count once (iota can be understood as the row index in the const statement block).
iota can be used as enumeration value.
*/

package main

import "fmt"

// enums
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const aa, bb, cc = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("The area is : %d", area)
	println()
	println(aa, bb, cc)

	// iota
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //independent value, iota += 1
		e        //"ha" iota += 1
		f = 100  //iota +=1
		g        //100 iota +=1
		h = iota //7, resume count
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

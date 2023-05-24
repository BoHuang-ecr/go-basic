/*
	&a get the address of memotry witch saves the value of variable a
	*p get the value of pointer p
*/

package main

import "fmt"

const MAX int = 3

func pointer() {
	var a int = 20 // define the actual variable
	var ip *int    // define the pointer variable

	ip = &a // save the address

	fmt.Printf("the address of variable a: %x\n", &a)

	fmt.Printf("pointer address: %x\n", ip)

	fmt.Printf("value of *ip: %d\n", *ip)
}

func pointer_array() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func pointer_to_pointer() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000
	ptr = &a
	pptr = &ptr

	fmt.Printf("variable a = %d\n", a)
	fmt.Printf("pointer variable *ptr = %d\n", *ptr)
	fmt.Printf("pointer to pointer variable **pptr = %d\n", **pptr)
}

func main() {
	pointer()
	pointer_array()
	pointer_to_pointer()
}

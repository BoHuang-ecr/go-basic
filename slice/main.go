/*
	Go language slices are an abstraction over arrays.

	The length of a Go array cannot be changed. Such a collection is not suitable in a specific scenario.
	Go provides a flexible and powerful built-in type slice ("dynamic array").
	Compared with an array, the length of a slice is not fixed. Yes, you can append elements, which may increase the capacity of the slice when appending.

	--> var identifier []type
	--> init:
		make([]Type, length, capacity)
		make([]Type, length)
		[]Type{}
		[]Type{value1, value2, ..., valueN}
*/

package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)

	var empty []int

	printSlice(empty)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*
	type_name(expression)
*/

package main

import (
	"fmt"
	"strconv"
)

func number() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean value: %f\n", mean)
}

func string_to_number() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("conversion error:", err)
	} else {
		fmt.Printf("string '%s' convert to int %d\n", str, num)
	}
}

func interface_conversion() {
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

func main() {
	number()
	string_to_number()
	interface_conversion()
}

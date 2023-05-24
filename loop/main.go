/*
	for init; condition; post { }
*/

package main

import "fmt"

func for_loop() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("index %d, value of x = %d\n", i, x)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// get key and value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// get key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// get value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}

func break_loop() {

	// break flag
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}

func continue_loop() {
	// continue flag
	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue re
		}
	}
}

func goto_loop() {
	var a int = 10

	// goto flag
LOOP:
	for a < 20 {
		if a == 15 {
			// jump out of loop
			a = a + 1
			goto LOOP
		}
		fmt.Printf("The value of a is: %d\n", a)
		a++
	}
}

func main() {
	for_loop()
	break_loop()
	continue_loop()
	goto_loop()
}

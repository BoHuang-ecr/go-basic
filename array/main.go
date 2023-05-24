/*
	var variable_name [SIZE] variable_type
*/

package main

import "fmt"

func init_array_1d() {
	var i, j, k int
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

func init_array_2d() {
	animals := [][]string{}

	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}

func main() {
	init_array_1d()
	init_array_2d()

	// use array as formal variable in functions
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32
	avg = getAverage(balance, 5)

	fmt.Printf("The average value is: %f ", avg)
}

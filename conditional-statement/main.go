/*
	if
	if else
	switch
	select
*/
package main

import (
	"fmt"
)

func if_else_func() {
	var a int = 100

	if a < 20 {
		fmt.Printf("a smaller than 20\n")
	} else {
		fmt.Printf("a bigger than 20\n")
	}
	fmt.Printf("a value: %d\n", a)

}

func if_nested() {
	var a int = 100
	var b int = 200

	if a == 100 {
		if b == 200 {
			fmt.Printf("The value of a is 100, the value of b is 200\n")
		}
	}
	fmt.Printf("a value: %d\n", a)
	fmt.Printf("b value: %d\n", b)
}

func swittch_func(x interface{}) {

	switch i := x.(type) {
	case nil:
		fmt.Printf("The tyoe of x is:%T\n", i)
	case int:
		fmt.Printf("x in type int\n")
	case float64:
		fmt.Printf("x is type float64\n")
	case func(int) float64:
		fmt.Printf("x is func(int)\n")
	case bool, string:
		fmt.Printf("x is bool or string\n")
	default:
		fmt.Printf("unknown\n")
	}
}

/*
Select:
	The select statement can only be used for channel operations, and each case must be a channel operation, either sending or receiving.
		Each case must be a channel
		All channel expressions are evaluated
		All expressions sent will be evaluated
		If any channel is available, it executes, the others are ignored.
		If there are multiple cases that can be run, select will randomly and fairly select one to execute, and the others will not be executed.
		otherwise:
			If there is a default clause, the statement is executed.
			Without the default clause, select blocks until a channel becomes available; Go does not re-evaluate channels or values.
*/

func select_func() {
	// define 2 channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	// start 2 goroutine，
	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	// use select to get data from channel
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// if there is no message from channel, run default
			fmt.Println("no message received")
		}
	}
}

func main() {
	if_else_func()
	if_nested()
	swittch_func(1)
	swittch_func(1.0)
	swittch_func(true)
	swittch_func("string")
	select_func()
}

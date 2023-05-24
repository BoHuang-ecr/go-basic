/*
	go function_name( parameters )

	channel:
		ch <- v    // send v to channel ch
		v := <-ch  // get v from channel ch, and pass the value to v
*/

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range will go over the channel,
	// once c finish sending 10 numbers, the channel will be closed
	for i := range c {
		fmt.Println(i)
	}
}

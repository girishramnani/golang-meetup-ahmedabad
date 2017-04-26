//Demonstrate the power of go routines
//Code sample for the talk by rob pike
package main

import (
	"fmt"
)

// START OMIT

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 2
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}

// END OMIT

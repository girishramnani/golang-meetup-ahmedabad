package main

import (
	"fmt"
	"time"
)

func NumberGenerator() <-chan int {
	outChan := make(chan int)
	go func() {
		for i := 0; ; i++ {
			outChan <- i
		}
	}()
	return outChan
}

func EvenGenerator(inchan <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for val := range inchan {
			if val&1 == 0 {
				outChan <- val
			}
		}
	}()
	return outChan
}

func merge(chs ...<-chan int) <-chan int {
	outChan := make(chan int)
	for _, ch := range chs {
		go func() {
			for value := range ch {
				outChan <- value
			}
		}()
	}
	return outChan
}

func main() {
	// Generate a common channel with input numbers
	gen := NumberGenerator()
	// Fan-out the channel to 2 go routines
	even1 := EvenGenerator(gen)
	even2 := EvenGenerator(gen)
	// Fan in the resulting even numbers
	mergeOut := merge(even1, even2)
	// provide a timeout to the stop the generation
	timeout := time.After(time.Millisecond * 5)
	for {
		select {
		case value := <-mergeOut:
			fmt.Println(value)
		case <-timeout:
			return
		}
	}

}

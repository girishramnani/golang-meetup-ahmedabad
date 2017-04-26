package main

import "time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

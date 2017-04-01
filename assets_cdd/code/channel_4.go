package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello" // lets feed it something
	}()
	fmt.Println(<-ch)
}

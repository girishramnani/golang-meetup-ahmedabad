package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "Hello" // lets feed it something

	fmt.Println(<-ch)
}

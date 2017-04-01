package main

import "fmt"

func main() {
	ch := make(chan string)

	ch <- "Hello" // lets feed it something

	fmt.Println(<-ch)
}

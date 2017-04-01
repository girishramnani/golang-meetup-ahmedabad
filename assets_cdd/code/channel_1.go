package main

import "fmt"

func main() {
	ch := make(chan string)

	fmt.Println(<-ch)
}

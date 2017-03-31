package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Go")
		wg.Done()
	}()

	wg.Wait()
}

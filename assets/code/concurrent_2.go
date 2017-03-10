package main

import "fmt"
import "time"
import "sync"

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)

	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("func 2")
		time.Sleep(1 * time.Second)

	}(&wg)

	wg.Wait()
	fmt.Println("Fin")
}

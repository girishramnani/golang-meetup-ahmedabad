package main 


import "sync"
import "time"
import "fmt"

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(5)

	go func(){

		for i := range int(5) {
			time.Sleep(1*time.Second)
			wg.Done()
		}
	}()

	select {
		case <- wg.Wait():
			fmt.Println("completed the wait")
	}

}
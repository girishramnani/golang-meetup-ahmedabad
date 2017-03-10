package main 

import (
    "fmt"
	"time"
)


func goroutine(i int) {
    fmt.Println("Waiting for",i,"ms")
    time.Sleep(time.Duration(i) * time.Millisecond)
}

func main() {

    go goroutine(200)
    go goroutine(300)
    go goroutine(400)

    fmt.Println("waiting for 1 second on main go routine")
    time.Sleep(1 * time.Second)
}
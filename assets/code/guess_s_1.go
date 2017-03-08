package main

import "fmt"

func main() {  
    x := []int{1,2,3}

    func(arr []int) {
        arr[0] = 7
        fmt.Println(arr) 
    }(x)

    fmt.Println(x) 
}
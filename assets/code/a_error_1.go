package main

import "fmt"

func doRecover() {  
    fmt.Println("recovered =>",recover()) 
}

func main() {  
    defer func() {
        doRecover() 
    }()

    panic("not good")
}

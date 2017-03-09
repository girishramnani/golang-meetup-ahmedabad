package main 

import (
	"fmt"
	"reflect"
)


func main() {
    var name string = "Hello"
    second := "World"

    var Unum uint8 = 12
    num := 65
    decinum := 56.433
    fmt.Println(name,second,Unum,decinum)
    fmt.Println(reflect.TypeOf(num))
}
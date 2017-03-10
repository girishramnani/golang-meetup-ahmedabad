package main 

import (
	"fmt"
)

func Map(function func(input int) int , inputs []int  ) []int {
    var output []int 
    for _,val  := range inputs {
        output = append(output,function(val))
    }
    return output
}
func main(){
    square := func(i int) int {
        return i*i
    }
    fmt.Println(Map(square,[]int {
        1,2,3,4,5,6,
    }))
}
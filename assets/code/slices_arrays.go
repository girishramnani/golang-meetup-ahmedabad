package main 

import (
	"fmt"
)


func main() {

    var myName [2]string 
    myName[0] = "Girish"
    myName[1] = "Ramnani"

    otherName  :=  [2]string {"Girish","Ramnani"}

    // slices with zeros

    mySlice := make([]int,5)
    otherSlice := []int{0,0,0,0,0}

    fmt.Println(myName,otherName,mySlice,otherSlice)
}
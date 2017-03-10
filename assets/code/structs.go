package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Firstname string 
	Lastname  string
	Email     string 
	Password  string 
}

func (u *User) FullName() string {
    return fmt.Sprint(u.Firstname," ",u.Lastname)
}

func (u User) Name() string {
    return fmt.Sprint(u.Firstname," ",u.Lastname)
}



func main() {
        u := User{"girish", "ramnani", "", "girish"}
        u2 := &User{"Girish","Ramnani","","Girish"} 
        u3 := new(User)
        fmt.Println(u,u.FullName())
        fmt.Println(u2,u2.FullName())
        fmt.Println(u3,u3.FullName())
        fmt.Println()
        b, _ := json.Marshal(u)
        os.Stdout.Write(b)
        
}
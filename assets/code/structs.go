package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User is a struct that will be used as orm mapping
type User struct {
	Firstname string `json:"first_name"`
	Lastname  string
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`
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
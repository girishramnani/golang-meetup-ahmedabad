package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Girish Ramnani"}
	t.Execute(os.Stdout, p)
}

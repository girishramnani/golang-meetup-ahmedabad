package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	templateString := `Golang HTTP Meetup`
	t, err := template.New("desc").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}

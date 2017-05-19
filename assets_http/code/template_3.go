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
	t, _ = t.Parse(`hello {{if .UserName -}}
					{{ .UserName }}!
					{{- else -}}
					Guest
					{{- end -}}`)
	p := Person{}
	t.Execute(os.Stdout, p)
}

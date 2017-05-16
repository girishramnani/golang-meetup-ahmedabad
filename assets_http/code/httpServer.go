package main

import (
	"net/http"
	"strings"
	"time"
)

type Time struct {
	format string
}

func (t Time) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format(t.format)))
}
func main() {
	th := Time{format: time.RFC1123}
	http.HandleFunc("/", IndexHandler)
	http.Handle("/time", th)
	http.ListenAndServe(":3000", nil)
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name != "" {
		name = strings.ToUpper(name)
		w.Write([]byte(name))
	} else {
		w.Write([]byte("Welcome "))
	}
}

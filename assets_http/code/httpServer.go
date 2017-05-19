package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

//START OMIT
type Time struct {
	format string
}

func (t Time) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format(t.format)))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	wr := bufio.NewWriter(w)
	wr.WriteString("Hello GoLang World")
	wr.Flush()
}

func main() {
	th := Time{format: time.RFC1123}
	http.HandleFunc("/", IndexHandler)
	http.Handle("/time", th)
	http.ListenAndServe(":3000", nil)
}

//END OMIT

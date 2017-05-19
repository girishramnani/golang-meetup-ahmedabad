package main

import (
	"fmt"
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func RequestMethod(method string) Adapter {

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there " + r.Method))
}

func main() {
	myHandler := http.HandlerFunc(IndexHandler)
	post := RequestMethod("POST")
	myHandlerPost := post(myHandler) // could be written as RequestMethod("POST")(myHandler)
	http.Handle("/", myHandlerPost)
	fmt.Println("Server started listening on :8081")
	http.ListenAndServe(":8081", nil)
}

package main

import "net/http"
import "net/http/httputil"
import "net/url"

type BasicAuth struct {
	Next     http.Handler
	username string
	password string
}

func NewBasicAuth(username, password string) *BasicAuth {
	return &BasicAuth{
		username: username,
		password: password,
	}
}
func (b *BasicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello world`))
}

func main() {
	userName := "girish"
	password := "girish"
	// get the username and password from the command line from the user

	_url, _ := url.Parse("https://github.com")
	proxy := httputil.NewSingleHostReverseProxy(_url)
	bscAuth := NewBasicAuth(userName, password)
	bscAuth.Next = proxy
	http.ListenAndServe(":8086", bscAuth)
}

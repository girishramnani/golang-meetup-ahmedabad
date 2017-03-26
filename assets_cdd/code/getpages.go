package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	links := []string{
		"http://usanewz.com/wings-of-hope/",
		"http://montrealgazette.com/sports/hockey/nhl/montreal-canadiens/stu-cowan-canadiens-andrew-shaw-and-artturi-lehkonen-make-a-good-pair",
		"https://bitbucket.org/inferno-os/inferno-os",
		"https://geth.ethereum.org/",
		"http://news.nationalgeographic.com/2017/03/human-heart-spinach-leaf-medicine-science",
	}

	var htmls []string

	for _, link := range links {
		data, _ := http.Get(link)
		html, _ := ioutil.ReadAll(data.Body)
		htmlstring := string(html)
		htmls = append(htmls, htmlstring)
	}

	fmt.Println(len(htmls))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// its a slice bro
func getHTML(link string, links []string, wg *sync.WaitGroup) {
	data, _ := http.Get(link)
	html, _ := ioutil.ReadAll(data.Body)
	links = append(links, string(html))
	wg.Done()
}

func main() {
	links := []string{
		"http://usanewz.com/wings-of-hope/",
		"http://montrealgazette.com/sports/hockey/nhl/montreal-canadiens/stu-cowan-canadiens-andrew-shaw-and-artturi-lehkonen-make-a-good-pair",
		"https://bitbucket.org/inferno-os/inferno-os",
		"https://geth.ethereum.org/",
		"http://news.nationalgeographic.com/2017/03/human-heart-spinach-leaf-medicine-science",
	}

	wg := &sync.WaitGroup{}
	var htmls []string

	for _, link := range links {
		wg.Add(1)
		go getHTML(link, htmls, wg)
	}
	wg.Wait()
	fmt.Println(len(htmls))
}

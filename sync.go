package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ch := make(chan string)

	go get("http://tkido.com/blog/", ch)
	go get("http://tkido.com/data/nicoime_latest.html", ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func get(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ch <- string(body)
}

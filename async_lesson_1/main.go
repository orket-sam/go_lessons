package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com/",
	"https://www.wikipedia.org/",
	"http://localhost:8081/",
	"http://localhost:8082/",
}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "cannot be reached")
		return
	}
	fmt.Println(url, "is up and running")

}

//	func main() {
//		for i := 0; i < len(urls); i++ {
//			go checkUrl(urls[i])
//		}
//		time.Sleep(5 * time.Second)
//	}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			checkUrl(url)
		}(urls[i])
	}

	wg.Wait()
}
